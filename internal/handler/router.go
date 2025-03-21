package handler

import (
	"github.com/mnkrana/crypto-balance/internal/adapters/commands"
	"github.com/mnkrana/crypto-balance/internal/ports"
	"github.com/mnkrana/crypto-balance/internal/utils"
	"github.com/spf13/cobra"
)

type Router struct {
	RpcPort  ports.RpcPort
	Commands map[string]ports.Command
}

func NewRouter(rpc ports.RpcPort) *Router {
	return &Router{
		RpcPort:  rpc,
		Commands: initializeCommands(rpc),
	}
}

func initializeCommands(rpc ports.RpcPort) map[string]ports.Command {
	return map[string]ports.Command{
		"balance":  commands.NewBalanceCommand(rpc),
		"token":    commands.NewTokenBalanceCommand(rpc),
		"transfer": commands.NewTransferCommand(rpc),
	}
}

func (r *Router) HandleRequest(action string, payload any) (string, error) {
	cmd, exists := r.Commands[action]
	if !exists {
		return utils.HandleError("Unknown command: "+action, nil)
	}
	return cmd.ExecuteRequest(action, payload)
}

func (r *Router) RegisterCommands() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "crypto-balance",
		Short: "Crypto balance CLI tool",
		Long:  `A CLI tool to interact with Ethereum wallets and tokens.`,
	}

	rootCmd.AddCommand(r.balanceCommand())
	rootCmd.AddCommand(r.tokenCommand())
	rootCmd.AddCommand(r.transferCommand())

	return rootCmd
}

func (r *Router) balanceCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "balance [wallet_address]",
		Short: "Get balance of a wallet",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			request := &commands.BalanceRequest{
				Address: args[0],
			}

			result, err := r.HandleRequest("balance", request)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Println(result)
		},
	}
}

func (r *Router) tokenCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "token [wallet_address] [token_address]",
		Short: "Get ERC-20 token balance",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			request := &commands.TokenRequest{
				Address: args[0],
				Token:   args[1],
			}

			result, err := r.HandleRequest("token", request)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Println(result)
		},
	}
}

func (r *Router) transferCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [to_address] [amount]",
		Short: "Transfer tokens",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			request := &commands.TransferRequest{
				From:   args[0],
				To:     args[1],
				Amount: args[2],
			}
			result, err := r.HandleRequest("transfer", request)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Println(result)
		},
	}
}
