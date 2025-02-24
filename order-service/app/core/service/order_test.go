package service

import (
	"context"
	"errors"
	mock_repository "github.com/huydq/order-service/app/core/mocks"
	"github.com/huydq/order-service/app/core/models"
	"github.com/huydq/order-service/util"
	pbOrderMgmt "github.com/huydq/proto/gen-go/order"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type OrderServiceTestSuite struct {
	suite.Suite
	mockRepo *mock_repository.MockIOrderRepository
	service  OrderService
}

func (suite *OrderServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.mockRepo = mock_repository.NewMockIOrderRepository(ctrl)
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
		mockResponse  *pbOrderMgmt.CreateOrderResponse
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
			mockResponse: &pbOrderMgmt.CreateOrderResponse{
				OrderId: int32(1),
			},
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
			mockResponse:  nil,
			mockError:     errors.New("create order failed"),
			expectedError: errors.New(util.ERR_INTERNAL_SERVER_ERROR),
		},
		{
			name:     "Create order with empty items",
			orderDto: pbOrderMgmt.CreateOrderRequest{CustomerId: 1, Items: []*pbOrderMgmt.OrderItemInput{}},
			mockResponse: &pbOrderMgmt.CreateOrderResponse{
				OrderId: int32(1),
			},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:     "Create order with invalid customer ID",
			orderDto: pbOrderMgmt.CreateOrderRequest{CustomerId: -1, Items: []*pbOrderMgmt.OrderItemInput{{ProductId: 1, Quantity: 2}}},
			mockResponse: &pbOrderMgmt.CreateOrderResponse{
				OrderId: int32(1),
			},
			mockError:     nil, // You might want to return an error here if invalid customer ID is an error condition
			expectedError: nil, // Update this based on your error handling
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
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

			if tc.mockError == nil {
				suite.mockRepo.EXPECT().CreateOrder(ctx, orderAgg).Return(int(tc.mockResponse.OrderId), nil)
			} else {
				suite.mockRepo.EXPECT().CreateOrder(ctx, orderAgg).Return(0, tc.mockError)
			}

			res, err := suite.service.CreateOrder(context.Background(), tc.orderDto)
			suite.Equal(tc.mockResponse, res)
			suite.Equal(tc.expectedError, err)
		})
	}
}
