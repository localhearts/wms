package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warehouse struct {
	WarehouseID   string         `gorm:"type:char(36);primaryKey;" json:"warehouse_id"`
	WarehouseCode string         `gorm:"not null;unique" json:"warehouse_code"`
	WarehouseName string         `gorm:"not null;" json:"warehouse_name"`
	Address       string         `gorm:"not null" json:"address"`
	CityID        string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"city_id"`
	ProvinceID    string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"province_id"`
	PostalCode    string         `gorm:"not null" json:"postal_code"`
	Phone         string         `gorm:"not null" json:"phone"`
	Capacity      int            `gorm:"not null" json:"capacity"`
	Manager       string         `gorm:"not null" json:"manager"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	City          City           `gorm:"references:CityID;" json:"city"`
	Province      Province       `gorm:"references:ProvinceID;" json:"province"`
	Storages      []Storage      `gorm:"foreignKey:WarehouseID;" json:"storages"`
}

// BeforeCreate untuk memastikan UUID digenerate jika belum ada
func (w *Warehouse) BeforeCreate(tx *gorm.DB) (err error) {
	if w.WarehouseID == "" {
		w.WarehouseID = uuid.New().String()
	}
	return
}
