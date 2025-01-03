package routes

import (
	"backend/controllers"
	"backend/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./course.go -destination=./mocks/course.go -package=mocks

type Storage interface {
	EnrollUserInCourse(userID string, courseID string) error
	GetCourse(courseID string) (*models.Course, error)
	GetCourses(userID string) ([]models.Course, error)
	SaveCourse(course models.Course) error
	UserKnowledge(userID string) (string, error)
}

type LLM interface {
	GenerateCourseBlueprint(courseHint string, userknoweldge string) (*models.CourseBlueprint, error)
}

type RouterSetup struct {
	controller controllers.Controller
}

func NewRouterSetup(storage Storage, llm LLM) RouterSetup {
	controller := controllers.NewController(storage, llm)
	return RouterSetup{controller: controller}
}

func (routerSetup RouterSetup) SetupCourseRoutes(router *gin.Engine) {

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
		courseGroup.GET("/generate", routerSetup.controller.GenerateCourse)
		courseGroup.POST("/enroll", routerSetup.controller.EnrollCourse)
		courseGroup.GET("/course", routerSetup.controller.GetCourse)
		courseGroup.GET("/courses", routerSetup.controller.GetCourses)
	}
}
