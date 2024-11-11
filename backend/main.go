package main

import (
	"backend/routes"
	"backend/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoStorage := storage.NewMustMongoDatabaseHandler()

	router := gin.Default()

	// Initialize routes
	routes.SetupCourseRoutes(router)

	// Start server
	router.Run(":8080")
}
