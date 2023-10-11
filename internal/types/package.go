package types

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	ProductID   string `json:"product_id"`
	PackageSize int    `json:"size"`
}
