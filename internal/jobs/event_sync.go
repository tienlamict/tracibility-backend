package jobs

import (
	"log"
	"tracibility/internal/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

func StartEventListener(db *gorm.DB, rpcURL string, contractAddress string) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Lỗi RPC: %v", err)
	}

	addr := common.HexToAddress(contractAddress)
	eth.ListenContractEvents(client, addr, db)
}
