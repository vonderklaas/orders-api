package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json: "order_id"`
	CustomerID  uuid.UUID  `json: "customer_id"`
	LineItems   []LineItem `json: "line_items"`
	CreatedAt   *time.Time `json: "created_id"`
	ShippedAt   *time.Time `json: "shipped_id"`
	CompletedAt *time.Time `json: "completed_id"`
}

type LineItem struct {
	ItemID   uuid.UUID `json: "item_id"`
	Quantity uint      `json: "quantity"`
	Price    uint      `json: "price"`
}
