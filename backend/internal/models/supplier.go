package models

import "gorm.io/gorm"

// Supplier represents a product provider
type Supplier struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:100;not null"`
	Email   string `json:"email" gorm:"size:100;unique"`
	Phone   string `json:"phone" gorm:"size:20"`
	Address string `json:"address" gorm:"size:200"`
}
