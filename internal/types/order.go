package types

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	Amount     int64  `json:"amount"`
}
