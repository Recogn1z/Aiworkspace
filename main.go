package main

import (
	"ai-workspace-backend/common"
	"ai-workspace-backend/router"
	"log"
)

func main() {
	if err := common.InitMYSQL(); err != nil {
		log.Printf("failed to initialize my sql: %v", err)
		return
	}

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Printf("failed to start server: %v", err)
		return
	}
}