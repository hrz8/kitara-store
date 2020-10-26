package models

import "github.com/gofrs/uuid"

// Order represents Order object for DB
type Order struct {
	ID          uuid.UUID        `gorm:"column:id;primaryKey" json:"id"`
	PriceAmount uint64           `gorm:"column:price_amount;min:1" json:"priceAmount"`
	Status      string           `gorm:"column:status;default:RESERVED" json:"status"`
	Products    []OrdersProducts `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"products"`
}
