package models

import "github.com/gofrs/uuid"

// Product represents an Product Item entity
type Product struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	Name     string    `gorm:"column:name;size:255;not null" json:"name"`
	Price    uint64    `gorm:"column:price;min:1" json:"price"`
	QtyTotal uint64    `gorm:"column:qty_total" json:"qtyTotal"`
}
