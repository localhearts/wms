package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	CategoryID   string         `gorm:"type:char(36);primaryKey;not null;" json:"category_id"`
	CategoryName string         `gorm:"not null;unique" json:"category_name"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft delete
}

// BeforeCreate ensures UUID is generated if it's empty
func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.CategoryID == "" {
		c.CategoryID = uuid.New().String()
	}
	return
}
