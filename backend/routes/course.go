package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCourseRoutes(router *gin.Engine) {
	courseGroup := router.Group("/course")
	{
		courseGroup.GET("/generate", controllers.GenerateCourse)
		courseGroup.POST("/enroll", controllers.EnrollCourse)
		courseGroup.GET("/get", controllers.GetCourse)
	}
}
