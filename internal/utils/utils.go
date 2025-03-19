package utils

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifyAddress(rawKey string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	isValid := re.MatchString(rawKey)
	return isValid
}

func VerifyKey(rawKey string) bool {
	re := regexp.MustCompile("^[0-9a-fA-F]{64}$")
	isValid := re.MatchString(rawKey)
	return isValid
}

func GetAddressFromRaw(raw string) (common.Address, error) {
	if !VerifyAddress(raw) {
		return common.Address{}, errors.New("address is not valid")
	}
	return common.HexToAddress(raw), nil
}

func GetPrivateKeyFromHex(rawPrivateKey string) (*ecdsa.PrivateKey, error) {
	if !VerifyKey(rawPrivateKey) {
		return nil, errors.New("private key is not valid")
	}
	privateKey, err := crypto.HexToECDSA(rawPrivateKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GetAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address
}

func ConvertToEther(amount *big.Int) *big.Float {
	famount := new(big.Float)
	famount.SetString(amount.String())
	ethValue := new(big.Float).Quo(famount, big.NewFloat(math.Pow10(18)))
	return ethValue
}

func ConvertToWei(amount *big.Float) *big.Int {
	fmt.Println("convert to wei", amount)
	ethValue := new(big.Float).Mul(amount, big.NewFloat(math.Pow10(18)))
	fmt.Println("converted", ethValue)
	bigIntValue, accuracy := ethValue.Int(nil)
	fmt.Println("Big Float Value:", ethValue)
	fmt.Println("Big Int Value:", bigIntValue)
	fmt.Println("Accuracy:", accuracy)
	fmt.Println("return", bigIntValue)
	return bigIntValue
}

func GetJsonResponse(input interface{}) string {
	e, err := json.Marshal(input)
	if err != nil {
		return err.Error()
	}
	return string(e)
}

func HandleError(message string, err error) (string, error) {
	if err != nil {
		log.Println(message, err)
	} else {
		log.Println(message)
	}
	return "", errors.New(message)
}
