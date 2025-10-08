package models

import "gorm.io/gorm"

// Sale represents a sale transaction

type Sale struct {
	gorm.Model
	ClientID  uint    `json:"client_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Total     float64 `json:"total" gorm:"not null"`
}
