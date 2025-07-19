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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

func ListenContractEvents(client *ethclient.Client, contractAddr common.Address, db *gorm.DB) {
	contractInstance, err := contracts.NewContracts(contractAddr, client)
	if err != nil {
		log.Fatalf("Bind contract error: %v", err)
	}

	// Tính topic hash cho từng sự kiện
	topicProductCreated := crypto.Keccak256Hash([]byte("ProductCreated(string,string,address)"))
	topicStepAdded := crypto.Keccak256Hash([]byte("StepAdded(string,string,string,address)"))

	// Lắng nghe tất cả log từ contract
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Subscribe error: %v", err)
	}

	log.Println("Lắng nghe sự kiện ProductCreated & StepAdded...")

	for {
		select {
		case err := <-sub.Err():
			log.Println("Subscription error:", err)

		case vLog := <-logs:
			if len(vLog.Topics) == 0 {
				continue
			}

			switch vLog.Topics[0] {
			case topicProductCreated:
				event, err := contractInstance.ParseProductCreated(vLog)
				if err != nil {
					log.Println("Parse ProductCreated lỗi:", err)
					continue
				}

				product := models.Product{
					ProductID:   event.ProductId,
					ProductName: event.Name,
					Location:    event.Location,
					Creator:     event.Creator.Hex(),
				}

				if err := db.Create(&product).Error; err != nil {
					log.Println("Ghi Product lỗi:", err)
				} else {
					log.Printf("Đã ghi ProductCreated: %s\n", event.ProductId)
				}

			case topicStepAdded:
				event, err := contractInstance.ParseStepAdded(vLog)
				if err != nil {
					log.Println("Parse StepAdded lỗi:", err)
					continue
				}

				statusString, ok := StepStatusMap[event.Status]
				if !ok {
					statusString = "Unknown"
				}

				trace := models.ProductTrace{
					TraceID:   uuid.New().String(),
					ProductID: event.ProductId,
					EventType: statusString,
					Location:  event.Location,
					Actor:     event.Actor.Hex(),
					TxHash:    vLog.TxHash.Hex(),
					TxStatus:  "SUCCESS",
				}

				if err := db.Create(&trace).Error; err != nil {
					log.Println("Ghi StepAdded lỗi:", err)
				} else {
					log.Printf("Đã ghi StepAdded: %s\n", event.ProductId)
				}
			default:
				log.Println("Không phải event đã xử lý, bỏ qua")
			}
		}
	}
}
