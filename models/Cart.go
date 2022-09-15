package models

import "gorm.io/gorm"

type Cart struct {
	//ID           string `json:"id" gorm:"primary_key"`
	gorm.Model
	ProductName string `json:"product_name"`
	PriceBox    string `json:"price_box"`
	PricePiece  string `json:"price_piece"`
	PriceDozen  string `json:"price_dozen"`
	BillingItem string `gorm:"billing_item"`
	Status      string `json:"status"`
}
