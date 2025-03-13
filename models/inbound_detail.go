package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	InboundStatusPending = iota + 1
	InboundStatusReceived
	InboundStatusPlaced
)

type InboundDetail struct {
	InboundDetailID string         `gorm:"type:char(36);primaryKey;" json:"inbound_detail_id"`
	InboundID       string         `gorm:"type:char(36);index;not null;constraint:OnDelete:CASCADE;" json:"inbound_id"`
	Status          int            `gorm:"default:1;" json:"status"`
	ProductID       string         `gorm:"type:char(36);index;not null;constraint:OnDelete:CASCADE;" json:"product_id"`
	Quantity        int            `gorm:"not null" json:"quantity"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Product Product `gorm:"references:ProductID;" json:"product"`
	Inbound Inbound `gorm:"references:InboundID;" json:"inbound"`
}

// GetStatus returns the human-readable status string
func (i *InboundDetail) GetStatus() string {
	switch i.Status {
	case InboundStatusReceived:
		return "Received"
	case InboundStatusPlaced:
		return "Placed"
	case InboundStatusPending:
		return "Pending"
	default:
		return "Unknown"
	}
}

// BeforeCreate hook to ensure defaults
func (i *InboundDetail) BeforeCreate(tx *gorm.DB) (err error) {
	if i.Status == 0 {
		i.Status = InboundStatusPending
	}
	return
}
