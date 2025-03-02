package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	csql "github.com/huydq/gokits/libs/storage/pg-client"
	"github.com/huydq/order-service/internal/core/domain/model"
	"github.com/huydq/order-service/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type OrderRepositoryIntegrationTestSuite struct {
	suite.Suite
	pgContainer     testcontainers.Container
	db              *sqlx.DB
	orderRepository *OrderRepository
	ctx             context.Context
}

func (suite *OrderRepositoryIntegrationTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	// Create a PostgreSQL container
	container, err := testcontainers.GenericContainer(suite.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:15",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     "testuser",
				"POSTGRES_PASSWORD": "testpassword",
				"POSTGRES_DB":       "testdb",
			},
			WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}
	suite.pgContainer = container

	// Get the mapped port
	mappedPort, err := container.MappedPort(suite.ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get mapped port: %v", err)
	}

	// Set the database connection string
	dbURL := fmt.Sprintf("postgres://testuser:testpassword@localhost:%d/testdb?sslmode=disable", mappedPort.Int())

	// Initialize the repository
	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	suite.db = db

	// Read and execute SQL script from file
	sqlScript, err := os.ReadFile("test/init.sql")
	if err != nil {
		log.Fatalf("Failed to read SQL script: %v", err)
	}

	_, err = db.Exec(string(sqlScript))
	if err != nil {
		log.Fatalf("Failed to setup test database: %v", err)
	}

	suite.orderRepository = NewOrderRepository(OrderPostgresClient{db: csql.BasePostgresSqlxDB{Client: db}})
}

func (suite *OrderRepositoryIntegrationTestSuite) TearDownSuite() {
	// Clean up tables after tests
	_, _ = suite.db.Exec("DROP TABLE order_items")
	_, _ = suite.db.Exec("DROP TABLE orders")

	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("Failed to terminate container: %v", err)
	}
}

func TestOrderRepositoryIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryIntegrationTestSuite))
}

func (suite *OrderRepositoryIntegrationTestSuite) TestCreateOrder() {
	testCases := []struct {
		name          string
		orderAgg      model.OrderAggregate
		expectError   bool
		expectedError string // Add this field
	}{
		{
			name: "Success - create order with single item",
			orderAgg: model.OrderAggregate{
				Order: model.Order{
					CustomerID:  1,
					TotalAmount: 100.50,
					Status:      util.ORDER_STATUS_PENDING,
				},
				Items: []model.OrderItem{
					{
						ProductID: 10,
						Quantity:  2,
						Price:     50.25,
					},
				},
			},
			expectError: false,
		},
		{
			name: "Success - create order with multiple items",
			orderAgg: model.OrderAggregate{
				Order: model.Order{
					CustomerID:  2,
					TotalAmount: 200.75,
					Status:      util.ORDER_STATUS_PENDING,
				},
				Items: []model.OrderItem{
					{
						ProductID: 11,
						Quantity:  1,
						Price:     100.00,
					},
					{
						ProductID: 12,
						Quantity:  3,
						Price:     33.58,
					},
				},
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			orderID, err := suite.orderRepository.CreateOrder(suite.ctx, tc.orderAgg)

			if tc.expectError {
				suite.Error(err)
				suite.Equal(0, orderID)
				if tc.expectedError != "" {
					suite.Contains(err.Error(), tc.expectedError)
				}
			} else {
				suite.NoError(err)
				suite.Greater(orderID, 0)

				// Verify order exists
				var order model.Order
				err := suite.db.Get(&order, "SELECT * FROM orders WHERE id = $1", orderID)
				suite.NoError(err)
				suite.Equal(tc.orderAgg.Order.CustomerID, order.CustomerID)
				suite.Equal(tc.orderAgg.Order.TotalAmount, order.TotalAmount)

				// Verify order items
				var items []model.OrderItem
				err = suite.db.Select(&items, "SELECT * FROM order_items WHERE order_id = $1", orderID)
				suite.NoError(err)
				suite.Len(items, len(tc.orderAgg.Items))
			}
		})
	}
}
