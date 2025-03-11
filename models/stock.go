package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stock struct {
	StockID     string         `gorm:"type:char(36);primaryKey" json:"stock_id"`
	ProductID   string         `gorm:"type:char(36);index;constraint:OnDelete:CASCADE;" json:"product_id"`
	WarehouseID string         `gorm:"type:char(36);index;constraint:OnDelete:CASCADE;" json:"warehouse_id"`
	StorageID   string         `gorm:"type:char(36);index;constraint:OnDelete:CASCADE;" json:"storage_id"`
	Quantity    int            `gorm:"default:0" json:"quantity"`
	LastUpdated time.Time      `gorm:"autoUpdateTime" json:"last_updated"`
	Product     Product        `gorm:"foreignKey:ProductID;" json:"product"`
	Warehouse   Warehouse      `gorm:"foreignKey:WarehouseID;" json:"warehouse"`
	Storage     Storage        `gorm:"foreignKey:StorageID;" json:"storage"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft delete
}

// BeforeCreate untuk memastikan UUID digenerate jika belum ada
func (s *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	if s.StockID == "" {
		s.StockID = uuid.New().String()
	}
	return
}
