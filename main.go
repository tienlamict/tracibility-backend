package main

import (
	"fmt"
	"tracibility/configs"
	"tracibility/internal/database"
	"tracibility/routes"
)

func main() {
	cfg := configs.LoadConfig()
	db := database.Connect(cfg.Database)

	r := routes.NewRouter(db)
	port := cfg.Server.Port
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	r.Run(":" + port)
}
