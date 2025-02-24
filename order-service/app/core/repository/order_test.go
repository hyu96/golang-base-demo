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
	"github.com/stretchr/testify/assert"
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
	defer mockDb.Close()
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
		name          string
		orderAgg      models.OrderAggregate
		mockQuery     string
		mockArgs      []driver.Value
		mockReturnID  int64
		mockReturnErr error
		expectedID    int
		expectedError error
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
			mockQuery:     regexp.QuoteMeta("INSERT INTO orders (customer_id, total_amount, status, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"),
			mockArgs:      []driver.Value{1, 100.0, util.ORDER_STATUS_PENDING},
			mockReturnID:  1,
			mockReturnErr: nil,
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Create order failed - order insert error",
			orderAgg: models.OrderAggregate{
				Order: models.Order{
					CustomerID:  1,
					TotalAmount: 100.0,
				},
				Items: []models.OrderItem{
					{ProductID: 1, Quantity: 2, Price: 50.0},
				},
			},
			mockQuery:     regexp.QuoteMeta("INSERT INTO orders (customer_id, total_amount, status, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"),
			mockArgs:      []driver.Value{1, 100.0, util.ORDER_STATUS_PENDING},
			mockReturnID:  0,
			mockReturnErr: errors.New("insert order error"),
			expectedID:    0,
			expectedError: errors.New("insert order error"),
		},
		// Add more test cases for different scenarios like empty items, item insert error, etc.
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.mockDB.ExpectBegin()
			orderRows := sqlmock.NewRows([]string{"id"}).AddRow(tc.mockReturnID)
			suite.mockDB.ExpectQuery(tc.mockQuery).WithArgs(tc.mockArgs...).WillReturnRows(orderRows).WillReturnError(tc.mockReturnErr)

			if tc.mockReturnErr == nil {
				// Only expect item insert if order insert was successful
				suite.mockDB.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
				)).
					WithArgs(tc.mockReturnID, tc.orderAgg.Items[0].ProductID, tc.orderAgg.Items[0].Quantity, tc.orderAgg.Items[0].Price).
					WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate successful insertion of 1 row
				suite.mockDB.ExpectCommit()
			} else {
				suite.mockDB.ExpectRollback()
			}

			orderID, err := suite.repository.CreateOrder(context.Background(), tc.orderAgg)

			assert.Equal(suite.T(), tc.expectedID, orderID)
			if tc.expectedError != nil {
				assert.Error(suite.T(), err)
				if tc.expectedError.Error() != util.ERR_INTERNAL_SERVER_ERROR {
					assert.Equal(suite.T(), tc.expectedError.Error(), err.Error())

				} else {
					assert.Equal(suite.T(), tc.expectedError, err)
				}
			} else {
				assert.NoError(suite.T(), err)
			}
		})
	}
}
