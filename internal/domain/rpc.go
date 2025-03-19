package domain

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Rpc struct {
	Network    string            `json:"network"`
	RpcUrl     string            `json:"rpc"`
	EthClient  *ethclient.Client `json:"client"`
	ChainId    *big.Int          `json:"chainId"`
	PrivateKey *ecdsa.PrivateKey `json:"pk"`
	Address    *common.Address   `json:"address"`
}
