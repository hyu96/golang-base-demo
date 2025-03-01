package model

type Order struct {
	ID          int     `db:"id" json:"id"`
	CustomerID  int     `db:"customer_id" json:"customer_id"`
	TotalAmount float64 `db:"total_amount" json:"total_amount"`
	Status      string  `db:"status" json:"status"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
	UpdatedAt   string  `db:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	ID        int     `db:"id" json:"id"`
	OrderID   int     `db:"order_id" json:"order_id"`
	ProductID int     `db:"product_id" json:"product_id"`
	Quantity  int     `db:"quantity" json:"quantity"`
	Price     float64 `db:"price" json:"price"`
}

type OrderAggregate struct {
	Order Order
	Items []OrderItem
}
