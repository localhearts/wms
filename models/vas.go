package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vas struct {
	VasID       string         `gorm:"type:char(36);primaryKey;" json:"vas_id"`
	ServiceName string         `gorm:"type:varchar(100);not null" json:"service_name"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:numeric(18,2);not null" json:"price"`
	InboundID   string         `gorm:"type:char(36);index;constraint:OnDelete:SET NULL;;" json:"inbound_id,omitempty"`
	OutboundID  string         `gorm:"type:char(36);index;constraint:OnDelete:SET NULL;;" json:"outbound_id,omitempty"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	// Hubungan dengan Inbound atau Outbound (optional)
	Inbound  *Inbound  `gorm:"foreignKey:InboundID;" json:"inbound,omitempty"`
	Outbound *Outbound `gorm:"foreignKey:OutboundID;" json:"outbound,omitempty"`
}

// Hook untuk auto generate UUID sebelum insert
func (v *Vas) BeforeCreate(tx *gorm.DB) (err error) {
	if v.VasID == "" {
		v.VasID = uuid.New().String()
	}
	return
}
