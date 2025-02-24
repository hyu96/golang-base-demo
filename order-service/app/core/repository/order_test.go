package repository

import (
	"context"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	csql "github.com/huydq/gokits/libs/storage/pg-client"
	"github.com/huydq/order-service/app/core/models"
	"github.com/huydq/order-service/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	mockDB     sqlmock.Sqlmock
	sqlxDB     *sqlx.DB
	repository OrderRepository
}

func (suite *OrderRepositoryTestSuite) SetupTest() {
	mockDb, mock, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDb, "sqlmock")
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	suite.mockDB = mock
	suite.sqlxDB = sqlxDB
	suite.repository = OrderRepository{orderClient: csql.BasePostgresSqlxDB{Client: suite.sqlxDB}}
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.sqlxDB.Close()
}

func TestOrderRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestCreateOrder() {
	testCases := []struct {
		name               string
		orderAgg           models.OrderAggregate
		mockOrderQuery     string
		mockOrderArgs      []driver.Value
		mockOrderReturnID  int64
		mockOrderReturnErr error
		mockItemsQuery     string
		mockItemsArgs      []driver.Value
		mockItemsErr       error
		expectedID         int
		expectedError      error
	}{
		{
			name: "Create order success",
			orderAgg: models.OrderAggregate{
				Order: models.Order{
					CustomerID:  1,
					TotalAmount: 100.0,
				},
				Items: []models.OrderItem{
					{ProductID: 1, Quantity: 2, Price: 50.0},
				},
			},
			mockOrderQuery:     regexp.QuoteMeta("INSERT INTO orders (customer_id, total_amount, status, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"),
			mockOrderArgs:      []driver.Value{1, 100.0, util.ORDER_STATUS_PENDING},
			mockOrderReturnID:  1,
			mockOrderReturnErr: nil,
			mockItemsQuery:     regexp.QuoteMeta("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)"),
			mockItemsArgs:      []driver.Value{1, 1, 2, 50.0},
			mockItemsErr:       nil,
			expectedID:         1,
			expectedError:      nil,
		},
		{
			name: "Create order failed",
			orderAgg: models.OrderAggregate{
				Order: models.Order{
					CustomerID:  1,
					TotalAmount: 100.0,
				},
				Items: []models.OrderItem{
					{ProductID: 1, Quantity: 2, Price: 50.0},
				},
			},
			mockOrderQuery:     regexp.QuoteMeta("INSERT INTO orders (customer_id, total_amount, status, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"),
			mockOrderArgs:      []driver.Value{1, 100.0, util.ORDER_STATUS_PENDING},
			mockOrderReturnID:  0,
			mockOrderReturnErr: errors.New("Can't create order"),
			mockItemsQuery:     regexp.QuoteMeta("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)"),
			mockItemsArgs:      []driver.Value{1, 1, 2, 50.0},
			mockItemsErr:       nil,
			expectedID:         0,
			expectedError:      errors.New(util.ERR_INTERNAL_SERVER_ERROR),
		},
		{
			name: "Create order item failed",
			orderAgg: models.OrderAggregate{
				Order: models.Order{
					CustomerID:  1,
					TotalAmount: 100.0,
				},
				Items: []models.OrderItem{
					{ProductID: 1, Quantity: 2, Price: 50.0},
				},
			},
			mockOrderQuery:     regexp.QuoteMeta("INSERT INTO orders (customer_id, total_amount, status, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"),
			mockOrderArgs:      []driver.Value{1, 100.0, util.ORDER_STATUS_PENDING},
			mockOrderReturnID:  1,
			mockOrderReturnErr: nil,
			mockItemsQuery:     regexp.QuoteMeta("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)"),
			mockItemsArgs:      []driver.Value{1, 1, 2, 50.0},
			mockItemsErr:       errors.New("Can't create order item"),
			expectedID:         0,
			expectedError:      errors.New(util.ERR_INTERNAL_SERVER_ERROR),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.mockDB.ExpectBegin()
			orderRows := sqlmock.NewRows([]string{"id"}).AddRow(tc.mockOrderReturnID)
			suite.mockDB.ExpectQuery(tc.mockOrderQuery).WithArgs(tc.mockOrderArgs...).WillReturnRows(orderRows).WillReturnError(tc.mockOrderReturnErr)

			if tc.mockOrderReturnErr != nil {
				suite.mockDB.ExpectRollback()
			} else {
				suite.mockDB.ExpectExec(tc.mockItemsQuery).WithArgs(tc.mockItemsArgs...).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.mockItemsErr)
				if tc.mockItemsErr != nil {
					suite.mockDB.ExpectRollback()
				} else {
					suite.mockDB.ExpectCommit()
				}
			}

			orderID, err := suite.repository.CreateOrder(context.Background(), tc.orderAgg)

			suite.Equal(tc.expectedID, orderID)
			if tc.expectedError != nil {
				suite.Error(err)
				suite.Equal(tc.expectedError.Error(), err.Error())
			} else {
				suite.NoError(err)
			}
		})
	}
}
