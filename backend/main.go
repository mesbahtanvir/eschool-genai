package main

import (
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	routes.SetupCourseRoutes(router)

	// Start server
	router.Run(":8080")
}
