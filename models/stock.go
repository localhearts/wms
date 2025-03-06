package models

import "time"

type Stock struct {
	StockID     uint `gorm:"primaryKey"`
	ProductID   uint
	WarehouseID uint
	LocationID  *uint
	Quantity    int       `gorm:"default:0"`
	LastUpdated time.Time `gorm:"autoUpdateTime"`
	Product     Product   `gorm:"foreignKey:ProductID"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID"`
	Location    Location  `gorm:"foreignKey:LocationID"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   time.Time `gorm:"index"`
}
