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
		log.Fatalf("Bind contract lỗi: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Lỗi subscribe: %v", err)
	}

	log.Println("⏳ Đang lắng nghe sự kiện StepAdded...")

	for {
		select {
		case err := <-sub.Err():
			log.Println("Lỗi subscription:", err)
		case vLog := <-logs:
			event, err := contractInstance.ParseStepAdded(vLog)
			if err != nil {
				log.Println("Parse log lỗi:", err)
				continue
			}

			trace := models.ProductTrace{
				ProductID: event.ProductId,
				EventType: event.EventType,
				Location:  event.Location,
				TxHash:    vLog.TxHash.Hex(),
				TxStatus:  "SUCCESS",
			}

			if err := db.Create(&trace).Error; err != nil {
				log.Println("Lỗi ghi DB:", err)
			} else {
				log.Printf("Event thêm bước truy xuất ghi DB: %s\n", event.ProductId)
			}
		}
	}
}
