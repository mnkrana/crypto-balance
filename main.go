package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/mnkrana/crypto-balance/internal/adapters"
	"github.com/mnkrana/crypto-balance/internal/handler"
)

func loadEnv() {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error retrieving executable path: %v", err)
	}
	execDir := filepath.Dir(execPath)
	envPath := filepath.Join(execDir, ".env")

	if err := godotenv.Overload(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
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
