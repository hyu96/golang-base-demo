package postgres

import (
	"context"
	"fmt"
	csql "github.com/huydq/gokits/libs/storage/pg-client"
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

var orderRepository *OrderRepository
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

	// Create tables and insert test data
	_, err = db.Exec(`
		CREATE TABLE orders (
			id SERIAL PRIMARY KEY,
			customer_id INT NOT NULL,
			total_amount FLOAT NOT NULL,
			status VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE order_items (
			id SERIAL PRIMARY KEY,
			order_id INT NOT NULL,
			product_id INT NOT NULL,
			quantity INT NOT NULL,
			price FLOAT NOT NULL,
			FOREIGN KEY (order_id) REFERENCES orders(id)
		);

		INSERT INTO orders (customer_id, total_amount, status) VALUES
		(1, 100.50, 'pending'),
		(2, 250.00, 'processing'),
		(3, 50.75, 'completed');

		INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
		(1, 10, 2, 50.25),
		(1, 11, 1, 0),
		(2, 20, 5, 50.00),
		(3, 30, 1, 50.75);
	`)
	if err != nil {
		log.Fatalf("Failed to setup test database: %v", err)
	}

	orderRepository = NewOrderRepository(OrderPostgresClient{db: csql.BasePostgresSqlxDB{Client: db}})

	// Run the tests
	code := m.Run()

	// Clean up tables after tests
	_, _ = db.Exec("DROP TABLE order_items")
	_, _ = db.Exec("DROP TABLE orders")

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
