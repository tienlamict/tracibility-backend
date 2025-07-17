package eth

import (
	"context"
	"log"
	"tracibility/internal/contracts"
	"tracibility/internal/database/models"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ListenToEvents(client *ethclient.Client, contractAddr common.Address, db *gorm.DB) {
	contractInstance, err := contracts.NewContracts(contractAddr, client)
	if err != nil {
		log.Fatalf("Failed to bind contract: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("subscribe error: %v", err)
	}

	log.Println("Listening for StepAdded events...")

	// Lắng nghe log
	for {
		select {
		case err := <-sub.Err():
			log.Println("Subscription error:", err)

		case vLog := <-logs:
			parsedLog, err := contractInstance.ParseStepAdded(vLog)
			if err != nil {
				log.Println("Failed to parse StepAdded event:", err)
				continue
			}

			trace := models.ProductTrace{
				ProductID: parsedLog.ProductId,
				EventType: parsedLog.EventType,
				Location:  parsedLog.Location,
				TxHash:    vLog.TxHash.Hex(),
				TxStatus:  "SUCCESS",
			}
			if err := db.Create(&trace).Error; err != nil {
				log.Println("DB insert error:", err)
			} else {
				log.Printf("StepAdded event saved: ProductID=%s\n", parsedLog.ProductId)
			}
		}
	}
}
