package eth

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTransactionStatus(client *ethclient.Client, txHash string) (string, error) {
	hash := common.HexToHash(txHash)
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		if err == ethereum.NotFound {
			return "PENDING", nil
		}
		return "", err
	}

	if receipt.Status == 1 {
		return "SUCCESS", nil
	}
	return "FAILED", nil
}
