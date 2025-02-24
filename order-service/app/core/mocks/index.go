// Code generated by MockGen. DO NOT EDIT.
// Source: index.go
//
// Generated by this command:
//
//	mockgen -source=index.go -destination=../mocks/index.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	models "github.com/huydq/order-service/app/core/models"
	gomock "go.uber.org/mock/gomock"
)

// MockIOrderRepository is a mock of IOrderRepository interface.
type MockIOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIOrderRepositoryMockRecorder
	isgomock struct{}
}

// MockIOrderRepositoryMockRecorder is the mock recorder for MockIOrderRepository.
type MockIOrderRepositoryMockRecorder struct {
	mock *MockIOrderRepository
}

// NewMockIOrderRepository creates a new mock instance.
func NewMockIOrderRepository(ctrl *gomock.Controller) *MockIOrderRepository {
	mock := &MockIOrderRepository{ctrl: ctrl}
	mock.recorder = &MockIOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOrderRepository) EXPECT() *MockIOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockIOrderRepository) CreateOrder(ctx context.Context, orderAgg models.OrderAggregate) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, orderAgg)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockIOrderRepositoryMockRecorder) CreateOrder(ctx, orderAgg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockIOrderRepository)(nil).CreateOrder), ctx, orderAgg)
}
