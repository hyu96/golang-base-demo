package dto

import "time"

// Request DTOs
type CreateOrderRequestDTO struct {
	CustomerID int                 `json:"customer_id"`
	Items      []OrderItemInputDTO `json:"items"`
}

type GetOrderRequestDTO struct {
	OrderID int `json:"order_id"`
}

// Response DTOs
type CreateOrderResponseDTO struct {
	OrderID int    `json:"order_id"`
	Message string `json:"message"`
}

// Order Data Transfer Objects
type OrderDTO struct {
	OrderID     int            `json:"order_id"`
	CustomerID  int            `json:"customer_id"`
	TotalAmount float64        `json:"total_amount"`
	Status      string         `json:"status"`
	Items       []OrderItemDTO `json:"items"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type OrderItemDTO struct {
	ItemID    int     `json:"item_id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// Input DTO for creating order items
type OrderItemInputDTO struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
