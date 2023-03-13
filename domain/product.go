package domain

import (
	"reflect"
	"time"
	"unicode"
)

type Entity[T any] struct {
	Fields T
}

type Product struct {
	ID            *int       `sql:"id" ignoreInsertUpdate:"true"`
	Name          *string    `sql:"name"`
	ImageURL      *string    `sql:"image_url"`
	Link          *string    `sql:"link"`
	LastBuyPrice  *float64   `sql:"last_buy_price"`
	LastSellPrice *float64   `sql:"last_sell_price"`
	CreatedAt     *time.Time `sql:"created_at"`
	UpdatedAt     *time.Time `sql:"updated_at"`
}

type PurchasedItem struct {
	ID            *int       `sql:"id" ignoreInsertUpdate:"true"`
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

type ProductList struct {
	Products []Product
	Next     bool
	Count    int64
}

func GetTableName(object interface{}) (tableName string) {
	typeOf := reflect.TypeOf(object)

	objectName := typeOf.Elem().Name()

	for i, v := range objectName {
		if unicode.IsUpper(v) && i != 0 {
			tableName += "_"
		}

		tableName += string(unicode.ToLower(v))
	}

	return
}
