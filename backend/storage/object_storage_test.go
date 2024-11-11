package storage

import (
	"backend/models"
	"backend/storage/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func InitWithMockDatabase(
	mockedCourseCollection MongoCollection,
	mockedUserCollection MongoCollection,
) {
	courseCollection = mockedCourseCollection
	userCollection = mockedCourseCollection
}
func TestSaveCourse(t *testing.T) {
	ctrl := gomock.NewController(t)

	testCases := []struct {
		name     string
		course   models.Course
		expected error
	}{
		{
			"When error from database Then return error",
			models.Course{
				CourseID: "123",
				CourseBlueprint: models.CourseBlueprint{
					Title:       "new course",
					Description: "new course description",
					Modules: []models.Module{
						{
							Title:       "module 1",
							Explanation: "this is module 1",
						},
					},
				},
			},
			nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SaveCourse(tc.course)
			cmp.Equal(
				result,
				tc.expected,
			)
		})
	}

	mockedUserCollection := mocks.NewMockMongoCollection(ctrl)
	mockedCourseCollection := mocks.NewMockMongoCollection(ctrl)
	InitWithMockDatabase(mockedUserCollection, mockedCourseCollection)

}
