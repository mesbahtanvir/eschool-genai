package controllers_test

import (
	"backend/controllers"
	"backend/controllers/mocks"
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGenerateCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		courseHint   string
		setupMocks   func(*mocks.MockStorage, *mocks.MockLLM)
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:       "Success",
			courseHint: "AI",
			setupMocks: func(mockStorage *mocks.MockStorage, mockLLM *mocks.MockLLM) {
				mockStorage.EXPECT().UserKnowledge(gomock.Any()).Return("Prior knowledge", nil)
				mockLLM.EXPECT().GenerateCourseBlueprint("AI", "Prior Knowledge").Return(&models.CourseBlueprint{Title: "Intro to AI"}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"course": map[string]interface{}{
					"course_blueprint": map[string]interface{}{
						"title":       "Intro to AI",
						"description": "",
						"modules":     interface{}(nil),
					},
					"course_id": "bf6de66d-e2c7-4a82-b348-5a9c36c95a99",
				},
			},
		},
		{
			name:       "LLM Failure",
			courseHint: "AI",
			setupMocks: func(mockStorage *mocks.MockStorage, mockLLM *mocks.MockLLM) {
				mockStorage.EXPECT().UserKnowledge(gomock.Any()).Return("Prior Knowledge", nil)
				mockLLM.EXPECT().GenerateCourseBlueprint("AI", "Prior Knowledge").Return(nil, errors.New("failed"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": "Failed to generate course blueprint",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLLM := mocks.NewMockLLM(ctrl)
			mockStorage := mocks.NewMockStorage(ctrl)
			tt.setupMocks(mockStorage, mockLLM)
			controller := controllers.NewController(mockStorage, mockLLM)

			router := gin.Default()
			router.GET("/generate_course", controller.GenerateCourse)

			req := httptest.NewRequest(http.MethodGet, "/generate_course?course_hint="+tt.courseHint, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedCode, resp.Code)

			var actualBody map[string]interface{}
			json.Unmarshal(resp.Body.Bytes(), &actualBody)

			actualCourse, ok := actualBody["course"].(map[string]interface{})
			if ok {
				delete(actualCourse, "course_id")
			}
			expectedCourse, ok := tt.expectedBody["course"].(map[string]interface{})
			if ok {
				delete(expectedCourse, "course_id")
			}

			assert.Equal(t, tt.expectedBody, actualBody)
		})
	}
}

func TestEnrollCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		userID       string
		courseID     string
		setupMocks   func(*mocks.MockStorage)
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:     "Success",
			userID:   "123",
			courseID: "456",
			setupMocks: func(mockStorage *mocks.MockStorage) {
				mockStorage.EXPECT().EnrollUserInCourse("123", "456").Return(nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"status": "success",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			mockLLM := mocks.NewMockLLM(ctrl)
			tt.setupMocks(mockStorage)

			controller := controllers.NewController(mockStorage, mockLLM)

			router := gin.Default()
			router.GET("/enroll_course", controller.EnrollCourse)

			req := httptest.NewRequest(http.MethodGet, "/enroll_course?user_id="+tt.userID+"&course_id="+tt.courseID, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedCode, resp.Code)

			var actualBody map[string]interface{}
			json.Unmarshal(resp.Body.Bytes(), &actualBody)
			assert.Equal(t, tt.expectedBody, actualBody)
		})
	}
}

func TestGetCourse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name         string
		courseID     string
		setupMocks   func(*mocks.MockStorage)
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:     "Success",
			courseID: "123",
			setupMocks: func(mockStorage *mocks.MockStorage) {
				mockStorage.EXPECT().GetCourse("123").Return(&models.Course{
					CourseID:        "123",
					CourseBlueprint: models.CourseBlueprint{Title: "Test Course"},
				}, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"course": map[string]interface{}{
					"course_blueprint": map[string]interface{}{
						"description": "",
						"modules":     interface{}(nil),
						"title":       "Test Course",
					},
					"course_id": "123"},
			},
		},
		{
			name:     "Not Found",
			courseID: "123",
			setupMocks: func(mockStorage *mocks.MockStorage) {
				mockStorage.EXPECT().GetCourse("123").Return(nil, mongo.ErrNoDocuments)
			},
			expectedCode: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error": "Course not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			mockLLM := mocks.NewMockLLM(ctrl)
			tt.setupMocks(mockStorage)

			controller := controllers.NewController(mockStorage, mockLLM)

			router := gin.Default()
			router.GET("/get_course", controller.GetCourse)

			req := httptest.NewRequest(http.MethodGet, "/get_course?course_id="+tt.courseID, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedCode, resp.Code)

			var actualBody map[string]interface{}
			json.Unmarshal(resp.Body.Bytes(), &actualBody)
			assert.Equal(t, tt.expectedBody, actualBody)
		})
	}
}
