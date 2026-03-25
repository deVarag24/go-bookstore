package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     uint    `gorm:"primaryKey" json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
