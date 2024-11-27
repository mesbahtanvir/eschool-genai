package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -source=./course_controller.go -destination=./mocks/course_controller.go -package=mocks

type Storage interface {
	EnrollUserInCourse(userID string, courseID string) error
	GetCourse(courseID string) (*models.Course, error)
	SaveCourse(course models.Course) error
	UserKnowledge(userID string) (string, error)
}

type LLM interface {
	GenerateCourseBlueprint(courseHint string, userKnowledge string) (*models.CourseBlueprint, error)
}

type Controller struct {
	storage Storage
	llm     LLM
}

func NewController(
	storage Storage,
	llm LLM,
) Controller {
	return Controller{
		storage: storage,
		llm:     llm,
	}
}

func (controller Controller) GenerateCourse(c *gin.Context) {
	courseHint := c.Query("course_hint")
	userID := c.Query("user_id")

	userKnowledge, err := controller.storage.UserKnowledge(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive user knowledge"})
		return
	}

	blueprint, err := controller.llm.GenerateCourseBlueprint(courseHint, userKnowledge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate course blueprint"})
		return
	}

	courseID := uuid.New().String()
	course := models.Course{
		CourseID:        courseID,
		CourseBlueprint: *blueprint,
	}

	// Save the course to the database
	if err := controller.storage.SaveCourse(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (controller Controller) EnrollCourse(c *gin.Context) {
	userID := c.Query("user_id")
	courseID := c.Query("course_id")

	// Append courseID to user's enrolled courses in object storage
	controller.storage.EnrollUserInCourse(userID, courseID)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (controller Controller) GetCourse(c *gin.Context) {
	courseID := c.Query("course_id")

	// Retrieve the course from storage
	course, err := controller.storage.GetCourse(courseID)
	if err != nil {
		// Check if the error is due to the course not being found
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		} else {
			// Handle other potential errors (e.g., database issues)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve course"})
		}
		return
	}

	// Respond with the retrieved course
	c.JSON(http.StatusOK, gin.H{"course": course})
}
