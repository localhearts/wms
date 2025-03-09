package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inbound struct {
	InboundID    string         `gorm:"type:char(36);primaryKey" json:"inbound_id"`
	ReferenceNo  string         `gorm:"not null;unique" json:"reference_no"`
	CustomerID   string         `gorm:"type:char(36);index" json:"customer_id"`
	InboundDate  time.Time      `gorm:"not null" json:"inbound_date"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ReceivedDate *time.Time     `json:"received_date,omitempty"`
	ReceivedBy   *string        `json:"received_by,omitempty"`
}

// BeforeCreate ensures UUIDs are generated if they are empty
func (i *Inbound) BeforeCreate(tx *gorm.DB) (err error) {
	if i.InboundID == "" {
		i.InboundID = uuid.New().String()
	}
	if i.CustomerID == "" {
		i.CustomerID = uuid.New().String()
	}
	return
}
