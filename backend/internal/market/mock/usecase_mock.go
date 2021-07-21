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
	utils "github.com/seregaa020292/capitalhub/pkg/utils"
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

// Create mocks base method.
func (m *MockUseCase) Create(ctx context.Context, asset *models.Market) (*models.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, asset)
	ret0, _ := ret[0].(*models.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(ctx, asset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), ctx, asset)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(ctx context.Context, assetID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, assetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(ctx, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), ctx, assetID)
}

// GetAll mocks base method.
func (m *MockUseCase) GetAll(ctx context.Context, pq *utils.PaginationQuery) (*models.MarketList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, pq)
	ret0, _ := ret[0].(*models.MarketList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUseCaseMockRecorder) GetAll(ctx, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUseCase)(nil).GetAll), ctx, pq)
}

// GetByID mocks base method.
func (m *MockUseCase) GetByID(ctx context.Context, assetID uuid.UUID) (*models.MarketBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, assetID)
	ret0, _ := ret[0].(*models.MarketBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUseCaseMockRecorder) GetByID(ctx, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUseCase)(nil).GetByID), ctx, assetID)
}

// SearchByTitle mocks base method.
func (m *MockUseCase) SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.MarketList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByTitle", ctx, title, query)
	ret0, _ := ret[0].(*models.MarketList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByTitle indicates an expected call of SearchByTitle.
func (mr *MockUseCaseMockRecorder) SearchByTitle(ctx, title, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByTitle", reflect.TypeOf((*MockUseCase)(nil).SearchByTitle), ctx, title, query)
}

// Update mocks base method.
func (m *MockUseCase) Update(ctx context.Context, asset *models.Market) (*models.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, asset)
	ret0, _ := ret[0].(*models.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUseCaseMockRecorder) Update(ctx, asset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), ctx, asset)
}
