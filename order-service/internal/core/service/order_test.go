package service

import (
	"context"
	"errors"
	"github.com/huydq/order-service/internal/core/domain/dto"
	models "github.com/huydq/order-service/internal/core/domain/model"
	mock_port "github.com/huydq/order-service/internal/core/port/mocks"

	mock_repository "github.com/huydq/order-service/internal/core/repository/mocks"
	"github.com/huydq/order-service/util"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type OrderServiceTestSuite struct {
	suite.Suite
	mockProductServiceClient *mock_port.MockIProductServiceClient
	mockRepo                 *mock_repository.MockIOrderRepository
	service                  OrderService
}

func (suite *OrderServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.mockRepo = mock_repository.NewMockIOrderRepository(ctrl)
	suite.mockProductServiceClient = mock_port.NewMockIProductServiceClient(ctrl)
	suite.service = OrderService{
		orderRepo:            suite.mockRepo,
		productServiceClient: suite.mockProductServiceClient,
	}
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (suite *OrderServiceTestSuite) TestCreateOrder() {
	type mockRepoData struct {
		input    models.OrderAggregate
		response *dto.CreateOrderResponseDTO
		err      error
	}

	type mockProductServiceData struct {
		response *dto.GetProductsResponseDTO
		err      error
	}

	testCases := []struct {
		name                   string
		orderDto               dto.CreateOrderRequestDTO
		mockRepoData           mockRepoData
		mockProductServiceData mockProductServiceData
		expectedResult         *dto.CreateOrderResponseDTO
		expectedError          error
	}{
		{
			name: "Create order success",
			orderDto: dto.CreateOrderRequestDTO{
				CustomerID: 1,
				Items: []dto.OrderItemInputDTO{
					{ProductID: 1, Quantity: 2},
				},
			},
			mockProductServiceData: mockProductServiceData{
				response: &dto.GetProductsResponseDTO{
					Products: []*dto.ProductDTO{
						{
							ID:    1,
							Name:  "Mock Data",
							Price: 100,
						},
					},
				},
			},
			mockRepoData: mockRepoData{
				input: models.OrderAggregate{
					Order: models.Order{
						CustomerID: 1,
					},
					Items: []models.OrderItem{
						{ProductID: 1, Quantity: 2, Price: 100},
					},
				},
				response: &dto.CreateOrderResponseDTO{
					OrderID: 1,
				},
				err: nil,
			},
			expectedResult: &dto.CreateOrderResponseDTO{OrderID: 1},
			expectedError:  nil,
		},
		{
			name: "Create order failed",
			orderDto: dto.CreateOrderRequestDTO{
				CustomerID: 1,
				Items: []dto.OrderItemInputDTO{
					{ProductID: 1, Quantity: 2},
				},
			},
			mockProductServiceData: mockProductServiceData{
				response: &dto.GetProductsResponseDTO{
					Products: []*dto.ProductDTO{
						{
							ID:    1,
							Name:  "Mock Data",
							Price: 100,
						},
					},
				},
			},
			mockRepoData: mockRepoData{
				input: models.OrderAggregate{
					Order: models.Order{
						CustomerID: 1,
					},
					Items: []models.OrderItem{
						{ProductID: 1, Quantity: 2, Price: 100},
					},
				},
				response: nil,
				err:      errors.New("create order failed"),
			},
			expectedResult: nil,
			expectedError:  errors.New(util.ERR_INTERNAL_SERVER_ERROR),
		},
		{
			name:     "Create order with empty items",
			orderDto: dto.CreateOrderRequestDTO{CustomerID: 1, Items: []dto.OrderItemInputDTO{}},
			mockProductServiceData: mockProductServiceData{
				response: &dto.GetProductsResponseDTO{
					Products: []*dto.ProductDTO{
						{
							ID:    1,
							Name:  "Mock Data",
							Price: 100,
						},
					},
				},
			},
			mockRepoData: mockRepoData{
				input: models.OrderAggregate{
					Order: models.Order{
						CustomerID: 1,
					},
					Items: []models.OrderItem{},
				},
				response: &dto.CreateOrderResponseDTO{
					OrderID: 1,
				},
				err: nil,
			},
			expectedResult: &dto.CreateOrderResponseDTO{OrderID: 1},
			expectedError:  nil,
		},
		{
			name: "Create order with invalid customer ID",
			orderDto: dto.CreateOrderRequestDTO{
				CustomerID: -1,
				Items: []dto.OrderItemInputDTO{
					{ProductID: 1, Quantity: 2},
				},
			},
			mockProductServiceData: mockProductServiceData{
				response: &dto.GetProductsResponseDTO{
					Products: []*dto.ProductDTO{
						{
							ID:    1,
							Name:  "Mock Data",
							Price: 100,
						},
					},
				},
			},
			mockRepoData: mockRepoData{
				input: models.OrderAggregate{
					Order: models.Order{
						CustomerID: -1,
					},
					Items: []models.OrderItem{
						{ProductID: 1, Quantity: 2, Price: 100},
					},
				},
				response: &dto.CreateOrderResponseDTO{
					OrderID: 1,
				},
				err: nil,
			},
			expectedResult: &dto.CreateOrderResponseDTO{OrderID: 1},
			expectedError:  nil, // Update this based on your error handling
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		suite.Run(tc.name, func() {
			productList := make([]int, len(tc.orderDto.Items))
			for index, item := range tc.orderDto.Items {
				productList[index] = item.ProductID
			}

			productReqDto := dto.GetProductsRequestDTO{
				ProductIDs: productList,
			}
			if tc.mockProductServiceData.err == nil {
				suite.mockProductServiceClient.EXPECT().GetProducts(ctx, productReqDto).Return(tc.mockProductServiceData.response, nil)
			} else {
				suite.mockProductServiceClient.EXPECT().GetProducts(ctx, productReqDto).Return(nil, tc.mockProductServiceData.err)
			}

			if tc.mockRepoData.err == nil {
				suite.mockRepo.EXPECT().CreateOrder(ctx, tc.mockRepoData.input).Return(tc.mockRepoData.response.OrderID, nil)
			} else {
				suite.mockRepo.EXPECT().CreateOrder(ctx, tc.mockRepoData.input).Return(0, tc.mockRepoData.err)
			}

			res, err := suite.service.CreateOrder(context.Background(), tc.orderDto)
			suite.Equal(tc.expectedResult, res)
			suite.Equal(tc.expectedError, err)
		})
	}
}
