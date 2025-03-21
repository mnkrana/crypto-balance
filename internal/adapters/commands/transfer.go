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

	to, err := utils.GetAddressFromRaw(req.To)
	if err != nil {
		return utils.HandleError("to address is invalid", err)
	}

	amount := new(big.Int)
	amount.SetString(req.Amount, 10)

	auth, err := m.Rpc.GetAuthByPrivateKey(pKey, amount)
	if err != nil {
		log.Fatal("failed to get auth:", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), to, amount, auth.GasLimit, auth.GasPrice, nil)

	chainID := big.NewInt(m.Rpc.GetChainId().Int64())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pKey)
	if err != nil {
		log.Fatal("failed to sign transaction:", err)
	}

	err = m.Rpc.GetEthClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("failed to send transaction:", err)
	}

	response := fmt.Sprintf("\033[32mTx sent: %s \033[0m", signedTx.Hash().Hex())
	return response, nil
}
