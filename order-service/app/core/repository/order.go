package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/huydq/order-service/app/core/models"
	"github.com/huydq/order-service/util"
	"github.com/jmoiron/sqlx"
	"strings"
)

func (o OrderRepository) CreateOrder(ctx context.Context, orderAgg models.OrderAggregate) (int, error) {
	var orderID int
	err := o.orderClient.Transaction(ctx, func(tx context.Context) error {
		orderArgs := []interface{}{
			orderAgg.Order.CustomerID,
			orderAgg.Order.TotalAmount,
			util.ORDER_STATUS_PENDING,
		}
		orderQuery := "INSERT INTO orders (customer_id, total_amount, status, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW()) RETURNING id"
		orderQuery = sqlx.Rebind(sqlx.DOLLAR, orderQuery)
		row := o.orderClient.GetTx(tx).QueryRowxContext(tx, orderQuery, orderArgs...)
		err := row.Scan(&orderID)
		if err != nil {
			log.Errorf("Create order failed", err.Error())
			return err
		}

		// Prepare batch insert for order items.
		valueStrings := []string{}
		itemValueArgs := []interface{}{}
		// Using 4 columns: order_id, product_id, quantity, price.
		for _, item := range orderAgg.Items {
			valueStrings = append(valueStrings, "(?, ?, ?, ?)")
			itemValueArgs = append(itemValueArgs, orderID, item.ProductID, item.Quantity, item.Price)
		}

		itemsQuery := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES %s", strings.Join(valueStrings, ","))
		itemsQuery = sqlx.Rebind(sqlx.DOLLAR, itemsQuery)
		_, err = o.orderClient.GetTx(tx).ExecContext(tx, itemsQuery, itemValueArgs...)
		if err != nil {
			log.Errorf("Create order item failed", err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		log.Errorf("Create order failed aggregate", err.Error())
		return 0, errors.New(util.ERR_INTERNAL_SERVER_ERROR)
	}

	return orderID, nil
}
