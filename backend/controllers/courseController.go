package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateCourse(c *gin.Context) {
	userID := c.Query("user_id")
	courseHint := c.Query("course_hint")

	blueprint, err := services.GenerateCourseBlueprint(courseHint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate course blueprint"})
		return
	}

	courseID := uuid.New().String()
	course := models.Course{
		CourseID:        courseID,
		CourseBlueprint: *blueprint,
	}

	// Save to object storage (logic to be added in storage package)

	c.JSON(http.StatusOK, course)
}

func EnrollCourse(c *gin.Context) {
	userID := c.Query("user_id")
	courseID := c.Query("course_id")

	// Append courseID to user's enrolled courses in object storage

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetCourse(c *gin.Context) {
	courseID := c.Query("course_id")

	// Retrieve course from object storage

	c.JSON(http.StatusOK, gin.H{"course": "course data here"})
}
