package models

import (
	"time"
)

type Supplier struct {
	SupplierID    uint      `gorm:"primaryKey"`
	SupplierName  string    `gorm:"type:varchar(255);not null"`
	ContactPerson string    `gorm:"type:varchar(255)"`
	Phone         string    `gorm:"type:varchar(20)"`
	Email         string    `gorm:"type:varchar(255)"`
	Address       string    `gorm:"type:varchar(255)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	DeletedAt     time.Time `gorm:"index"`
}
