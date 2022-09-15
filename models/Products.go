package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	PriceBox    string `json:"price_box"`
	PricePiece  string `json:"price_piece"`
	PriceDozen  string `json:"price_dozen"`
	Status      string `json:"status"`
}
