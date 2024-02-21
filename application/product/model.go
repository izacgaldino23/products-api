package product

import "time"

type Product struct {
	ID            *int64     `json:"id"`
	Name          *string    `json:"name" validate:"required"`
	ImageURL      *string    `json:"image_url" validate:"required"`
	Link          *string    `json:"link" validate:"required"`
	LastBuyPrice  *float64   `json:"last_buy_price"`
	LastSellPrice *float64   `json:"last_sell_price"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	// RemovedAt     *time.Time `json:"removed_at"`
}

type ProductList struct {
	Products []Product `json:"products"`
	Next     bool      `json:"next"`
	Count    int64     `json:"count"`
}
