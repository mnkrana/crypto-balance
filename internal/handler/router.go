package handler

import (
	"github.com/mnkrana/crypto-balance/internal/adapters/commands"
	"github.com/mnkrana/crypto-balance/internal/ports"
	"github.com/mnkrana/crypto-balance/internal/utils"
)

type Router struct {
	RpcPort  ports.RpcPort
	Commands map[string]ports.Command
}

func NewRouter(
	rpc ports.RpcPort,
) *Router {
	return &Router{
		RpcPort:  rpc,
		Commands: initializeCommands(rpc),
	}
}

func initializeCommands(
	rpc ports.RpcPort,
) map[string]ports.Command {

	commands := map[string]ports.Command{
		"balance": commands.NewBalanceCommand(rpc),
		"token":   commands.NewTokenBalanceCommand(rpc),
	}
	return commands
}

func (r *Router) HandleRequest(action string, payload any) (string, error) {
	cmd, exists := r.Commands[action]
	if !exists {
		return utils.HandleError("Unknown command: "+action, nil)
	}

	return cmd.ExecuteRequest(action, payload)
}
