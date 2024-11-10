package controllers

import (
	"backend/models"
	"backend/services"
	"backend/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateCourse(c *gin.Context) {
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

	// Save the course to the database
	if err := storage.SaveCourse(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save course"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func EnrollCourse(c *gin.Context) {
	userID := c.Query("user_id")
	courseID := c.Query("course_id")

	// Append courseID to user's enrolled courses in object storage
	storage.EnrollUserInCourse(userID, courseID)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetCourse(c *gin.Context) {
	courseID := c.Query("course_id")

	// Retrieve the course from storage
	course, err := storage.GetCourse(courseID)
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
