// Code generated by MockGen. DO NOT EDIT.
// Source: pg_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	"github.com/seregaa020292/capitalhub/internal/asset/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	utils "github.com/seregaa020292/capitalhub/pkg/utils"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, asset *model.Asset) (*model.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, asset)
	ret0, _ := ret[0].(*model.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, asset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, asset)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, assetID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, assetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, assetID)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll(ctx context.Context, userID uuid.UUID) (*[]model.AssetBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, userID)
	ret0, _ := ret[0].(*[]model.AssetBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), ctx, userID)
}

// GetAllByMarketID mocks base method.
func (m *MockRepository) GetAllByMarketID(ctx context.Context, marketID uuid.UUID, query *utils.PaginationQuery) (*model.AssetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByMarketID", ctx, marketID, query)
	ret0, _ := ret[0].(*model.AssetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByMarketID indicates an expected call of GetAllByMarketID.
func (mr *MockRepositoryMockRecorder) GetAllByMarketID(ctx, marketID, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByMarketID", reflect.TypeOf((*MockRepository)(nil).GetAllByMarketID), ctx, marketID, query)
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(ctx context.Context, assetID uuid.UUID) (*model.AssetBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, assetID)
	ret0, _ := ret[0].(*model.AssetBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(ctx, assetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), ctx, assetID)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, asset *model.Asset) (*model.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, asset)
	ret0, _ := ret[0].(*model.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, asset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, asset)
}
