package main

import (
	"fmt"
	"tracibility/configs"
	"tracibility/internal/database"
	"tracibility/internal/eth"
	"tracibility/internal/jobs"
	"tracibility/routes"
)

func main() {
	cfg := configs.LoadConfig()
	db := database.Connect(cfg.Database)
	contractCfg := eth.EthereumConfig{
		RPCUrl:          cfg.Ethereum.RPCUrl,
		PrivateKey:      cfg.Ethereum.PrivateKey,
		ContractAddress: cfg.Ethereum.ContractAddress,
	}
	contractService, err := eth.NewContractService(contractCfg)
	if err != nil {
		fmt.Printf("Không thể khởi tạo contract service: %v", err)
	}

	go jobs.StartEventListener(db, contractService.Config.RPCUrl, contractService.Config.ContractAddress)

	r := routes.NewRouter(db)
	port := cfg.Server.Port
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	r.Run(":" + port)
}
