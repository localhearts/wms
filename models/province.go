package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Province struct {
	ProvinceID   string         `gorm:"type:char(36);primaryKey;" json:"province_id"`
	ProvinceName string         `gorm:"not null;unique" json:"province_name"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft delete
}

// BeforeCreate ensures a UUID is generated if it's empty
func (c *Province) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ProvinceID == "" {
		c.ProvinceID = uuid.New().String()
	}
	return
}
