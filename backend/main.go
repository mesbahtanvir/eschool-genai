package main

import (
	"backend/routes"
	"backend/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := storage.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	router := gin.Default()

	// Initialize routes
	routes.SetupCourseRoutes(router)

	// Start server
	router.Run(":8080")
}
