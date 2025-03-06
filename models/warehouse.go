package models

import "time"

type Warehouse struct {
	WarehouseID   uint      `gorm:"primaryKey"`
	WarehouseName string    `gorm:"type:varchar(255);not null"`
	Location      string    `gorm:"type:varchar(255)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	DeletedAt     time.Time `gorm:"index"`
}
