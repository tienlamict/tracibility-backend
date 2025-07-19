package eth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"tracibility/internal/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractService struct {
	Client   *ethclient.Client
	Contract *contracts.Contracts
	Config   EthereumConfig // chứa private key, contract address
}

type EthereumConfig struct {
	RPCUrl          string
	PrivateKey      string
	ContractAddress string
}

// Khởi tạo contract service (được gọi từ main.go)
func NewContractService(cfg EthereumConfig) (*ContractService, error) {
	client, err := ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		return nil, err
	}

	contractAddr := common.HexToAddress(cfg.ContractAddress)
	contract, err := contracts.NewContracts(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &ContractService{
		Client:   client,
		Contract: contract,
		Config:   cfg,
	}, nil
}

func HashObjectSHA256(obj interface{}) (string, error) {
	// Convert object to JSON bytes (consistent string representation)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(jsonBytes)
	return hex.EncodeToString(hash[:]), nil
}
