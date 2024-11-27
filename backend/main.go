package main

import (
	"backend/routes"
	"backend/services"
	"backend/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	dbStorage := storage.NewMustMongoDatabaseHandler()
	llm := services.NewOpenAIService()

	router := gin.Default()

	// Initialize routes
	routerSetup := routes.NewRouterSetup(dbStorage, llm)
	routerSetup.SetupCourseRoutes(router)

	// Start server
	router.Run(":8080")
}
