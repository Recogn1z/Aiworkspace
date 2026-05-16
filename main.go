package main

import (
	"ai-workspace-backend/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}