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
	InboundDetailID uint           `gorm:"primaryKey;autoIncrement" json:"inbound_detail_id"`
	ReferenceNo     string         `gorm:"not null;unique" json:"reference_no"`
	InboundID       string         `gorm:"type:char(36);index;not null" json:"inbound_id"`
	Status          int            `gorm:"default:1;index" json:"status"`
	Inbound         Inbound        `gorm:"foreignKey:InboundID;references:InboundID;constraint:OnDelete:CASCADE;" json:"inbound"`
	ProductID       string         `gorm:"type:char(36);index;not null" json:"product_id"`
	Product         Product        `gorm:"foreignKey:ProductID;references:ProductID;constraint:OnDelete:SET NULL;" json:"product"`
	Quantity        int            `gorm:"not null" json:"quantity"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
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
