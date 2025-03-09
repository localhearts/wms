package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	CityID    string         `gorm:"type:char(36);primaryKey;" json:"city_id"`
	CityName  string         `gorm:"not null;unique" json:"city_name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// BeforeCreate hook to generate UUID manually
func (c *City) BeforeCreate(tx *gorm.DB) (err error) {
	if c.CityID == "" {
		c.CityID = uuid.New().String()
	}
	return
}
