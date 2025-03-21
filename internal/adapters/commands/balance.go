package commands

import (
	"context"
	"fmt"
	"math/big"

	"github.com/mnkrana/crypto-balance/internal/ports"
	"github.com/mnkrana/crypto-balance/internal/utils"
)

type BalanceRequest struct {
	Address string
}

type BalanceCommand struct {
	Rpc ports.RpcPort
}

func NewBalanceCommand(port ports.RpcPort) *BalanceCommand {
	return &BalanceCommand{Rpc: port}
}

func (m *BalanceCommand) ExecuteRequest(action string, request any) (string, error) {
	req, ok := request.(*BalanceRequest)
	if !ok {
		return utils.HandleError("invalid request format", nil)
	}

	if req.Address == "" {
		return utils.HandleError("address is required", nil)
	}

	address, err := utils.GetAddressFromRaw(req.Address)
	if err != nil {
		return utils.HandleError("address is invalid", err)
	}

	balance, err := m.Rpc.GetEthClient().BalanceAt(context.Background(), address, nil)
	if err != nil {
		return utils.HandleError("balance command failed", err)
	}

	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	response := fmt.Sprintf("\033[32mBalance: %f ETH\033[0m", ethValue)
	return response, nil
}
