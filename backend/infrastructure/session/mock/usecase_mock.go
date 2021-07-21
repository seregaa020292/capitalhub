// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	models "github.com/seregaa020292/capitalhub/internal/models"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CleanMaxSession mocks base method.
func (m *MockUseCase) CleanMaxSession(ctx context.Context, userID uuid.UUID) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanMaxSession", ctx, userID)
	ret0, _ := ret[0].(int64)
	return ret0
}

// CleanMaxSession indicates an expected call of CleanMaxSession.
func (mr *MockUseCaseMockRecorder) CleanMaxSession(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanMaxSession", reflect.TypeOf((*MockUseCase)(nil).CleanMaxSession), ctx, userID)
}

// CreateSession mocks base method.
func (m *MockUseCase) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, session, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockUseCaseMockRecorder) CreateSession(ctx, session, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockUseCase)(nil).CreateSession), ctx, session, expire)
}

// DeleteByID mocks base method.
func (m *MockUseCase) DeleteByID(ctx context.Context, userID uuid.UUID, sessionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, userID, sessionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockUseCaseMockRecorder) DeleteByID(ctx, userID, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockUseCase)(nil).DeleteByID), ctx, userID, sessionID)
}

// GetSessionByID mocks base method.
func (m *MockUseCase) GetSessionByID(ctx context.Context, userID uuid.UUID, sessionID string) (*models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByID", ctx, userID, sessionID)
	ret0, _ := ret[0].(*models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByID indicates an expected call of GetSessionByID.
func (mr *MockUseCaseMockRecorder) GetSessionByID(ctx, userID, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByID", reflect.TypeOf((*MockUseCase)(nil).GetSessionByID), ctx, userID, sessionID)
}

// RefreshByID mocks base method.
func (m *MockUseCase) RefreshByID(ctx context.Context, session *models.Session, newSessionID string, expire int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshByID", ctx, session, newSessionID, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshByID indicates an expected call of RefreshByID.
func (mr *MockUseCaseMockRecorder) RefreshByID(ctx, session, newSessionID, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshByID", reflect.TypeOf((*MockUseCase)(nil).RefreshByID), ctx, session, newSessionID, expire)
}
