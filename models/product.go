package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ProductID   string         `gorm:"type:char(36);primaryKey;" json:"product_id"`
	SupplierID  string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"supplier_id"`
	UomID       string         `gorm:"type:char(36);not null;index" json:"uom_id"`
	CategoryID  string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"category_id"`
	ProductName string         `gorm:"not null;unique" json:"product_name"`
	SKU         string         `gorm:"not null;unique" json:"sku"`
	ExpiryDate  *time.Time     `json:"expiry_date,omitempty"`
	Length      float64        `gorm:"not null" json:"length"`
	Width       float64        `gorm:"not null" json:"width"`
	Height      float64        `gorm:"not null" json:"height"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Uom         Uom            `gorm:"references:UomID;" json:"uoms"`
	Category    Category       `gorm:"references:CategoryID;" json:"categories"`
}

// BeforeCreate hook to ensure UUID is generated
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ProductID == "" {
		p.ProductID = uuid.New().String()
	}
	return
}
