package commands

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mnkrana/crypto-balance/internal/ports"
	"github.com/mnkrana/crypto-balance/internal/utils"
)

type TransferRequest struct {
	From   string //private key
	To     string //address
	Amount string
}

type TransferCommand struct {
	Rpc ports.RpcPort
}

func NewTransferCommand(port ports.RpcPort) *TransferCommand {
	return &TransferCommand{Rpc: port}
}

func (m *TransferCommand) ExecuteRequest(action string, request any) (string, error) {
	req, ok := request.(*TransferRequest)
	if !ok {
		return utils.HandleError("invalid request format", nil)
	}

	if req.From == "" || req.To == "" || req.Amount == "" {
		return utils.HandleError("invalid request", nil)
	}

	pKey, err := utils.GetPrivateKeyFromHex(req.From)
	if err != nil {
		return utils.HandleError("private key is invalid", err)
	}

	fromAddress := utils.GetAddressFromPrivateKey(pKey)

	to, err := utils.GetAddressFromRaw(req.To)
	if err != nil {
		return utils.HandleError("to address is invalid", err)
	}

	nonce, err := m.Rpc.GetEthClient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("failed to fetch nonce:", err)
	}

	amount := new(big.Int)
	amount.SetString(req.Amount, 10)

	gasPrice, err := m.Rpc.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("failed to get gas price:", err)
	}

	gasLimit := uint64(21000)

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	chainID := big.NewInt(m.Rpc.GetChainId().Int64())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pKey)
	if err != nil {
		log.Fatal("failed to sign transaction:", err)
	}

	err = m.Rpc.GetEthClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("failed to send transaction:", err)
	}

	fmt.Println("Transaction sent! Hash:", signedTx.Hash().Hex())

	return signedTx.Hash().Hex(), nil
}
