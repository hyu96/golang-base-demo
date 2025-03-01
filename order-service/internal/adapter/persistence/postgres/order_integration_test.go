package postgres_test

import (
	"context"
	"fmt"
	"github.com/huydq/order-service/internal/adapter/persistence/postgres"
	"github.com/huydq/order-service/internal/core/domain/model"
	"github.com/huydq/order-service/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"testing"
	"time"
)

var orderRepository *postgres.OrderRepository
var db *sqlx.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Create a PostgreSQL container
	pgContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
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
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("Failed to terminate container: %v", err)
		}
	}()

	// Get the mapped port
	mappedPort, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get mapped port: %v", err)
	}

	// Set the database connection string
	dbURL := fmt.Sprintf("postgres://testuser:testpassword@localhost:%d/testdb?sslmode=disable", mappedPort.Int())

	// Initialize the repository
	db, err = sqlx.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	orderRepository = postgres.NewOrderRepository(postgres.OrderPostgresClient{db: postgres.BasePostgresSqlxDB{Client: db}})

	// Run the tests
	code := m.Run()

	os.Exit(code)
}

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()
	orderAgg := model.OrderAggregate{
		Order: model.Order{
			CustomerID:  1,
			TotalAmount: 100.50,
			Status:      util.ORDER_STATUS_PENDING,
		},
		Items: []model.OrderItem{
			{
				OrderID:   1,
				ProductID: 10,
				Quantity:  2,
				Price:     50.25,
			},
		},
	}

	orderID, err := orderRepository.CreateOrder(ctx, orderAgg)
	assert.NoError(t, err)
	assert.Greater(t, orderID, 0)
}
