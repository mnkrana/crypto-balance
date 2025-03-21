package main

import (
	"log"

	"github.com/mnkrana/crypto-balance/internal/adapters"
	"github.com/mnkrana/crypto-balance/internal/handler"
)

func initDependencies() *handler.Router {
	chainConfig := adapters.NewChain()
	rpcAdapter := adapters.NewRPCAdapter(chainConfig)
	rpcAdapter.Print()
	return handler.NewRouter(rpcAdapter)
}

func main() {
	handler.LoadEnv()
	router := initDependencies()

	if err := router.RegisterCommands().Execute(); err != nil {
		log.Fatal(err)
	}
}
