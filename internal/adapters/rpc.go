package adapters

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mnkrana/crypto-balance/internal/domain"
	"github.com/mnkrana/crypto-balance/internal/utils"
)

type RPC struct {
	Rpc *domain.Rpc
}

func NewRPCAdapter(chain *domain.Chain) *RPC {
	rpc := &RPC{}
	client, err := ethclient.DialContext(context.Background(), chain.RpcUrl)
	if err != nil {
		log.Fatalf("failed to initialize RPC adapter: %v", err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to initialize RPC adapter: %v", err)
	}

	privateKey, err := utils.GetPrivateKeyFromHex(chain.PKey)
	if err != nil {
		log.Fatalf("failed to initialize RPC adapter: %v", err)
	}

	address := utils.GetAddressFromPrivateKey(privateKey)

	rpc.Rpc = &domain.Rpc{
		Network:    chain.Network,
		RpcUrl:     chain.RpcUrl,
		EthClient:  client,
		ChainId:    chainId,
		PrivateKey: privateKey,
		Address:    &address,
	}

	return rpc
}

func (svc *RPC) GetEthClient() *ethclient.Client {
	return svc.Rpc.EthClient
}

func (svc *RPC) GetChainId() *big.Int {
	return svc.Rpc.ChainId
}

func (svc *RPC) GetUserAddress() common.Address {
	return *svc.Rpc.Address
}
func (svc *RPC) GetPK() *ecdsa.PrivateKey {
	return svc.Rpc.PrivateKey
}

func (svc *RPC) GetAuth(value *big.Int) (*bind.TransactOpts, error) {
	address := utils.GetAddressFromPrivateKey(svc.Rpc.PrivateKey)
	nonce, err := svc.Rpc.EthClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal("error in geting Nonce", err)
	}

	gasPrice, err := svc.Rpc.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("error in getting gas price", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(svc.Rpc.PrivateKey, svc.Rpc.ChainId)
	if err != nil {
		fmt.Println("error in getting key transactor", err)
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = 21000
	auth.GasPrice = big.NewInt(0).Mul(gasPrice, big.NewInt(2))
	return auth, nil
}

func (svc *RPC) GetAuthByPrivateKey(pKey *ecdsa.PrivateKey, value *big.Int) (*bind.TransactOpts, error) {
	address := utils.GetAddressFromPrivateKey(pKey)
	nonce, err := svc.Rpc.EthClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal("error in geting Nonce", err)
	}

	gasPrice, err := svc.Rpc.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("error in getting gas price", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pKey, svc.Rpc.ChainId)
	if err != nil {
		fmt.Println("error in getting key transactor", err)
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = 21000
	auth.GasPrice = gasPrice
	return auth, nil
}

func (svc *RPC) Print() {
	fmt.Println("Network: ", svc.Rpc.Network)
}

func (svc *RPC) GetDetails() string {
	e, err := json.Marshal(svc.Rpc)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(e)
}
