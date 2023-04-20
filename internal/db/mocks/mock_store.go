// Code generated by MockGen. DO NOT EDIT.
// Source: db.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	model "papercup-test/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAnnotation mocks base method.
func (m *MockStore) CreateAnnotation(newAnnotaton *model.Annotation) (*model.Annotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnnotation", newAnnotaton)
	ret0, _ := ret[0].(*model.Annotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAnnotation indicates an expected call of CreateAnnotation.
func (mr *MockStoreMockRecorder) CreateAnnotation(newAnnotaton interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnnotation", reflect.TypeOf((*MockStore)(nil).CreateAnnotation), newAnnotaton)
}

// CreateVideo mocks base method.
func (m *MockStore) CreateVideo(newVideo *model.Video) (*model.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVideo", newVideo)
	ret0, _ := ret[0].(*model.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVideo indicates an expected call of CreateVideo.
func (mr *MockStoreMockRecorder) CreateVideo(newVideo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVideo", reflect.TypeOf((*MockStore)(nil).CreateVideo), newVideo)
}

// DeleteAnnotation mocks base method.
func (m *MockStore) DeleteAnnotation(annotationID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnnotation", annotationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnnotation indicates an expected call of DeleteAnnotation.
func (mr *MockStoreMockRecorder) DeleteAnnotation(annotationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnnotation", reflect.TypeOf((*MockStore)(nil).DeleteAnnotation), annotationID)
}

// DeleteVideo mocks base method.
func (m *MockStore) DeleteVideo(videoID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVideo", videoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVideo indicates an expected call of DeleteVideo.
func (mr *MockStoreMockRecorder) DeleteVideo(videoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVideo", reflect.TypeOf((*MockStore)(nil).DeleteVideo), videoID)
}

// GetAnnotation mocks base method.
func (m *MockStore) GetAnnotation(id uuid.UUID) (*model.Annotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotation", id)
	ret0, _ := ret[0].(*model.Annotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnnotation indicates an expected call of GetAnnotation.
func (mr *MockStoreMockRecorder) GetAnnotation(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotation", reflect.TypeOf((*MockStore)(nil).GetAnnotation), id)
}

// GetAnnotationVideoID mocks base method.
func (m *MockStore) GetAnnotationVideoID(annotationID uuid.UUID) (*uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotationVideoID", annotationID)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnnotationVideoID indicates an expected call of GetAnnotationVideoID.
func (mr *MockStoreMockRecorder) GetAnnotationVideoID(annotationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotationVideoID", reflect.TypeOf((*MockStore)(nil).GetAnnotationVideoID), annotationID)
}

// GetUserID mocks base method.
func (m *MockStore) GetUserID(username, password string) (*uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", username, password)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserID indicates an expected call of GetUserID.
func (mr *MockStoreMockRecorder) GetUserID(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockStore)(nil).GetUserID), username, password)
}

// GetVideoLength mocks base method.
func (m *MockStore) GetVideoLength(videoID uuid.UUID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideoLength", videoID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideoLength indicates an expected call of GetVideoLength.
func (mr *MockStoreMockRecorder) GetVideoLength(videoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideoLength", reflect.TypeOf((*MockStore)(nil).GetVideoLength), videoID)
}

// ListAnnotations mocks base method.
func (m *MockStore) ListAnnotations(videoID uuid.UUID) ([]*model.Annotation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAnnotations", videoID)
	ret0, _ := ret[0].([]*model.Annotation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAnnotations indicates an expected call of ListAnnotations.
func (mr *MockStoreMockRecorder) ListAnnotations(videoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnnotations", reflect.TypeOf((*MockStore)(nil).ListAnnotations), videoID)
}

// UpdateAnnotation mocks base method.
func (m *MockStore) UpdateAnnotation(annotationID uuid.UUID, updatedAnno *model.AnnotationMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnnotation", annotationID, updatedAnno)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAnnotation indicates an expected call of UpdateAnnotation.
func (mr *MockStoreMockRecorder) UpdateAnnotation(annotationID, updatedAnno interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnnotation", reflect.TypeOf((*MockStore)(nil).UpdateAnnotation), annotationID, updatedAnno)
}
