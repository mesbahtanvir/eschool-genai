package routes

import (
	"backend/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type MongoStorage interface {
}

type RouterSetup struct {
	storage MongoStorage
}

func NewRouterSetup(storage MongoStorage) RouterSetup {
	return RouterSetup{storage: storage}
}

func (RouterSetup) SetupCourseRoutes(router *gin.Engine) {

	// Configure CORS settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	courseGroup := router.Group("/course")
	{
		courseGroup.GET("/generate", controllers.GenerateCourse)
		courseGroup.POST("/enroll", controllers.EnrollCourse)
		courseGroup.GET("/get", controllers.GetCourse)
	}
}
