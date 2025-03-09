package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	CustomerID string         `gorm:"primaryKey;type:char(36)" json:"customer_id"`
	FullName   string         `gorm:"not null" json:"full_name"`
	Phone      string         `gorm:"not null" json:"phone"`
	Address    string         `gorm:"not null" json:"address"`
	CityID     string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"city_id"`
	ProvinceID string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"province_id"`
	PostalCode string         `gorm:"not null" json:"postal_code"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	City       City           `gorm:"foreignKey:CityID" json:"city"`
	Province   Province       `gorm:"foreignKey:ProvinceID" json:"province"`
}

// BeforeCreate ensures UUID is generated if it's empty
func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	if c.CustomerID == "" {
		c.CustomerID = uuid.New().String()
	}
	return
}
