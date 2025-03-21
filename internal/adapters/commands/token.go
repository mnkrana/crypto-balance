package commands

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mnkrana/crypto-balance/internal/ports"
	"github.com/mnkrana/crypto-balance/internal/utils"
)

type TokenRequest struct {
	Address string
	Token   string
}

type TokenBalanceCommand struct {
	Rpc ports.RpcPort
}

func NewTokenBalanceCommand(port ports.RpcPort) *TokenBalanceCommand {
	return &TokenBalanceCommand{Rpc: port}
}

func (m *TokenBalanceCommand) ExecuteRequest(action string, request any) (string, error) {
	req, ok := request.(*TokenRequest)
	if !ok {
		return utils.HandleError("invalid request format", nil)
	}

	if req.Address == "" || req.Token == "" {
		return utils.HandleError("invalid request", nil)
	}

	address, err := utils.GetAddressFromRaw(req.Address)
	if err != nil {
		return utils.HandleError("address is invalid", err)
	}

	token, err := utils.GetAddressFromRaw(req.Token)
	if err != nil {
		return utils.HandleError("token is invalid", err)
	}

	//get token balance without ABI
	balanceOfSelector := "70a08231"
	walletPadded := common.LeftPadBytes(address.Bytes(), 32)

	data := append(common.FromHex(balanceOfSelector), walletPadded...)

	callMsg := ethereum.CallMsg{
		To:   &token,
		Data: data,
	}

	result, err := m.Rpc.GetEthClient().CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return utils.HandleError("balance command failed", err)
	}

	balance := new(big.Int).SetBytes(result)
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	response := fmt.Sprintf("\033[32mBalance: %f ETH\033[0m", ethValue)
	return response, nil
}
