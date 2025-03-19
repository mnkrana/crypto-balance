package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mnkrana/crypto-balance/internal/adapters"
	"github.com/mnkrana/crypto-balance/internal/handler"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error: .env file not loaded: %v", err)
	}
}

func initDependencies() *adapters.RPC {
	chainConfig := adapters.NewChain()
	rpcAdapter := adapters.NewRPCAdapter(chainConfig)
	rpcAdapter.Print()
	return rpcAdapter
}

func main() {
	loadEnv()

	rpcAdapter := initDependencies()
	router := handler.NewRouter(rpcAdapter)
	router.HandleRequest("balance", os.Args[1])
}
