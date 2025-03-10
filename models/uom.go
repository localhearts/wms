package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Uom struct {
	UomID     string         `gorm:"type:char(36);primaryKey;not null;" json:"uom_id"`
	UomName   string         `gorm:"not null;unique" json:"uom_name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft delete
}

func (Uom) TableName() string {
	return "uoms"
}

// BeforeCreate ensures UUID is generated if it's empty
func (u *Uom) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UomID == "" {
		u.UomID = uuid.New().String()
	}
	return
}
