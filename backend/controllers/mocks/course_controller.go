// Code generated by MockGen. DO NOT EDIT.
// Source: ./course_controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "backend/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// EnrollUserInCourse mocks base method.
func (m *MockStorage) EnrollUserInCourse(userID, courseID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnrollUserInCourse", userID, courseID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnrollUserInCourse indicates an expected call of EnrollUserInCourse.
func (mr *MockStorageMockRecorder) EnrollUserInCourse(userID, courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnrollUserInCourse", reflect.TypeOf((*MockStorage)(nil).EnrollUserInCourse), userID, courseID)
}

// GetCourse mocks base method.
func (m *MockStorage) GetCourse(courseID string) (*models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourse", courseID)
	ret0, _ := ret[0].(*models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourse indicates an expected call of GetCourse.
func (mr *MockStorageMockRecorder) GetCourse(courseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourse", reflect.TypeOf((*MockStorage)(nil).GetCourse), courseID)
}

// GetCourses mocks base method.
func (m *MockStorage) GetCourses(couseID string) ([]models.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourses", couseID)
	ret0, _ := ret[0].([]models.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourses indicates an expected call of GetCourses.
func (mr *MockStorageMockRecorder) GetCourses(couseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourses", reflect.TypeOf((*MockStorage)(nil).GetCourses), couseID)
}

// SaveCourse mocks base method.
func (m *MockStorage) SaveCourse(course models.Course) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCourse", course)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCourse indicates an expected call of SaveCourse.
func (mr *MockStorageMockRecorder) SaveCourse(course interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCourse", reflect.TypeOf((*MockStorage)(nil).SaveCourse), course)
}

// UserKnowledge mocks base method.
func (m *MockStorage) UserKnowledge(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserKnowledge", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserKnowledge indicates an expected call of UserKnowledge.
func (mr *MockStorageMockRecorder) UserKnowledge(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserKnowledge", reflect.TypeOf((*MockStorage)(nil).UserKnowledge), userID)
}

// MockLLM is a mock of LLM interface.
type MockLLM struct {
	ctrl     *gomock.Controller
	recorder *MockLLMMockRecorder
}

// MockLLMMockRecorder is the mock recorder for MockLLM.
type MockLLMMockRecorder struct {
	mock *MockLLM
}

// NewMockLLM creates a new mock instance.
func NewMockLLM(ctrl *gomock.Controller) *MockLLM {
	mock := &MockLLM{ctrl: ctrl}
	mock.recorder = &MockLLMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLLM) EXPECT() *MockLLMMockRecorder {
	return m.recorder
}

// GenerateCourseBlueprint mocks base method.
func (m *MockLLM) GenerateCourseBlueprint(courseHint, userKnowledge string) (*models.CourseBlueprint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCourseBlueprint", courseHint, userKnowledge)
	ret0, _ := ret[0].(*models.CourseBlueprint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCourseBlueprint indicates an expected call of GenerateCourseBlueprint.
func (mr *MockLLMMockRecorder) GenerateCourseBlueprint(courseHint, userKnowledge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCourseBlueprint", reflect.TypeOf((*MockLLM)(nil).GenerateCourseBlueprint), courseHint, userKnowledge)
}
