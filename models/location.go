package models

import "time"

type Location struct {
	LocationID   uint `gorm:"primaryKey"`
	WarehouseID  uint
	LocationCode string    `gorm:"type:varchar(50);unique"`
	Zone         string    `gorm:"type:varchar(50)"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	Warehouse    Warehouse `gorm:"foreignKey:WarehouseID"`
}
