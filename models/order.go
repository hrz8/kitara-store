package models

import "github.com/gofrs/uuid"

// Order represents Order object for DB
type Order struct {
	ID          uuid.UUID        `gorm:"column:id;primaryKey" json:"id"`
	PriceAmount uint64           `gorm:"column:price_amount;min:1" json:"priceAmount"`
	Products    []OrdersProducts `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"products"`
}

// CreateOrderPayload represent Order payload for creating Order
type CreateOrderPayload struct {
	Products []CreateOrdersProductsPayload `json:"products" validate:"required"`
}

// CreateOrdersProductsPayload represent Order Product payload for creating Order
type CreateOrdersProductsPayload struct {
	ID  uuid.UUID `json:"id" validate:"required"`
	Qty uint64    `json:"qty" validate:"required,gt=0"`
}
