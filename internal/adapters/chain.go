package adapters

import (
	"os"

	"github.com/mnkrana/crypto-balance/internal/domain"
	"github.com/mnkrana/crypto-balance/internal/utils"
)

func NewChain() *domain.Chain {
	network := os.Getenv(utils.NetworkKey)
	chain := &domain.Chain{
		Network: os.Getenv(network + utils.ChainNetworkKey),
		RpcUrl:  os.Getenv(network + utils.ChainRPCUrlKey),
		PKey:    os.Getenv(network + utils.ChainPKey),
	}
	return chain
}
