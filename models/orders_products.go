package models

import "github.com/gofrs/uuid"

// OrdersProducts represents an many to many Order<->Product
type OrdersProducts struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"column:order_id;not null;index" json:"orderId"`
	ProductID uuid.UUID `gorm:"not null;index" json:"productId"`
	Qty       uint64    `gorm:"column:qty;min:1" json:"qty"`
	Product   Product   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product"`
}
