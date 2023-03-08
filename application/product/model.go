package product

import "time"

type Product struct {
	ID            *int       `json:"id"`
	Name          *string    `json:"name" validate:"required"`
	ImageURL      *string    `json:"image_url" validate:"required"`
	Link          *string    `json:"link" validate:"required"`
	LastBuyPrice  *float64   `json:"last_buy_price"`
	LastSellPrice *float64   `json:"last_sell_price"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
