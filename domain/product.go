package domain

import "time"

type Product struct {
	ID            *int       `sql:"id" ignoreInsertUpdate:"true"`
	Name          *string    `sql:"name"`
	LastBuyPrice  *float64   `sql:"last_buy_price"`
	LastSellPrice *float64   `sql:"last_sell_price"`
	CreatedAt     *time.Time `sql:"created_at"`
	UpdatedAt     *time.Time `sql:"updated_at"`
}

type PurchasedItem struct {
	ID            *int       `sql:"id" ignoreInsertUpdate:"true"`
	Link          *string    `sql:"link"`
	ImageURL      *string    `sql:"image_url"`
	ProductID     *int       `sql:"product_id"`
	BuyPrice      *float64   `sql:"buy_price"`
	Amount        *int       `sql:"amount"`
	SubItemAmount *int       `sql:"sub_item_amount"`
	CreatedAt     *time.Time `sql:"created_at"`
	UpdatedAt     *time.Time `sql:"updated_at"`
}

type Purchase struct {
	ID          *int       `sql:"id" ignoreInsertUpdate:"true"`
	Supplier    *string    `sql:"supplier"`
	Description *string    `sql:"description"`
	CreatedAt   *time.Time `sql:"created_at"`
	UpdatedAt   *time.Time `sql:"updated_at"`
	Items       []PurchasedItem
}

type SoldItem struct {
	ID        *int       `sql:"id" ignoreInsertUpdate:"true"`
	ProductID *int       `sql:"product_id"`
	SellPrice *float64   `sql:"sell_price"`
	Amount    *int       `sql:"amount"`
	Taxa      *float64   `sql:"taxa"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type Sold struct {
	ID          *int       `sql:"id" ignoreInsertUpdate:"true"`
	PlatformID  *int       `sql:"platform_id"`
	Description *string    `sql:"description"`
	CreatedAt   *time.Time `sql:"created_at"`
	UpdatedAt   *time.Time `sql:"updated_at"`
	Items       []SoldItem
}
