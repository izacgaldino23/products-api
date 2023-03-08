package product

import "time"

type Product struct {
	ID            *int       `json:"id"`
	Name          *string    `json:"name" binding:"required"`
	ImageURL      *string    `json:"image_url" binding:"required"`
	Link          *string    `json:"link" binding:"required"`
	LastBuyPrice  *float64   `json:"last_buy_price"`
	LastSellPrice *float64   `json:"last_sell_price"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
