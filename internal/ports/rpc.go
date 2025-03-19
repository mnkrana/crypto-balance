package ports

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RpcPort interface {
	GetEthClient() *ethclient.Client
	GetChainId() *big.Int
	GetUserAddress() common.Address
	GetPK() *ecdsa.PrivateKey
	GetAuth(value *big.Int) (*bind.TransactOpts, error)
	Print()
	GetDetails() string
}
