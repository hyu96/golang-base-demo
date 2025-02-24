package service
package service

import (
	"context"
	"errors"
	"github.com/huydq/order-service/app/core/models"
	"github.com/huydq/order-service/mocks"
	pbOrderMgmt "github.com/huydq/proto/gen-go/order"
	"github.com/huydq/order-service/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OrderServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.IOrderRepository
	service  OrderService
}

func (suite *OrderServiceTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.IOrderRepository)
	suite.service = OrderService{
		orderRepo: suite.mockRepo,
	}
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (suite *OrderServiceTestSuite) TestCreateOrder() {
	testCases := []struct {
		name          string
		orderDto      pbOrderMgmt.CreateOrderRequest
		mockResponse  int
		mockError     error
		expectedError error
	}{
		{
			name: "Create order success",
			orderDto: pbOrderMgmt.CreateOrderRequest{
				CustomerId: 1,
				Items: []*pbOrderMgmt.OrderItemInput{
					{ProductId: 1, Quantity: 2},
				},
			},
			mockResponse:  1,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name: "Create order failed",
			orderDto: pbOrderMgmt.CreateOrderRequest{
				CustomerId: 1,
				Items: []*pbOrderMgmt.OrderItemInput{
					{ProductId: 1, Quantity: 2},
				},
			},
			mockResponse:  0,
			mockError:     errors.New("create order failed"),
			expectedError: errors.New(util.ERR_INTERNAL_SERVER_ERROR),
		},
		{
			name:          "Create order with empty items",
			orderDto:      pbOrderMgmt.CreateOrderRequest{CustomerId: 1, Items: []*pbOrderMgmt.OrderItemInput{}},
			mockResponse:  0,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:     "Create order with invalid customer ID",
			orderDto: pbOrderMgmt.CreateOrderRequest{CustomerId: -1, Items: []*pbOrderMgmt.OrderItemInput{{ProductId: 1, Quantity: 2}}},
			mockResponse:  0,
			mockError:     nil, // You might want to return an error here if invalid customer ID is an error condition
			expectedError: nil, // Update this based on your error handling
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			orderAgg := models.OrderAggregate{
				Order: models.Order{
					CustomerID: int(tc.orderDto.CustomerId),
				},
				Items: make([]models.OrderItem, len(tc.orderDto.Items)),
			}
			for i, item := range tc.orderDto.Items {
				orderAgg.Items[i] = models.OrderItem{
					ProductID: int(item.ProductId),
					Quantity:  int(item.Quantity),
				}
			}

			suite.mockRepo.On("CreateOrder", mock.Anything, orderAgg).Return(tc.mockResponse, tc.mockError)

			res, err := suite.service.CreateOrder(context.Background(), tc.orderDto)

			if tc.expectedError != nil {
				assert.Error(suite.T(), err)
				assert.Equal(suite.T(), tc.expectedError, err)
				assert.Nil(suite.T(), res)
			} else {
				assert.NoError(suite.T(), err)
				if tc.mockResponse != 0 {
					assert.Equal(suite.T(), int32(tc.mockResponse), res.OrderId)
				} else {
					// Add assertions for cases where mockResponse is 0 but no error is expected
					// For example, check if the order ID is 0 or some other default value
				}
			}
		})
	}
}
