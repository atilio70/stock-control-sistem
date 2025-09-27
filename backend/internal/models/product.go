package models

import "gorm.io/gorm"

//Product represents a product in the stock
type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
