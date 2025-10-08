package models

import "gorm.io/gorm"

//Product represents a product in the stock system
type Product struct {
	gorm.Model         //includes ID, CreatedAt, UpdateAt, DeleteAt
	Name       string  `json:"name" gorm:"size:100;not null"`
	SKU        string  `json:"sku" gorm:"size:50;unique;not null"`
	Price      float64 `json:"price" gorm:"not null"`
	Quantity   int     `json:"quantity" gorm:"not null"`
}
