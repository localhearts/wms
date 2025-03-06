package models

import (
	"time"
)

type Product struct {
	ProductID   uint   `gorm:"primaryKey"`
	ProductName string `gorm:"type:varchar(255);not null"`
	SKU         string `gorm:"type:varchar(50);unique"`
	Category    string `gorm:"type:varchar(100)"`
	SupplierID  *uint
	Unit        string    `gorm:"type:varchar(50)"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Supplier    Supplier  `gorm:"foreignKey:SupplierID"`
}
