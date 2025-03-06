package models

import "time"

type Inbound struct {
	InboundID           uint `gorm:"primaryKey"`
	CustomerID          uint
	InboundNotification time.Time `gorm:"autoCreateTime"`
	Status              string    `gorm:"type:enum('Pending','Processing','Completed');default:'Pending'"`
	CreatedAt           time.Time `gorm:"autoCreateTime"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime"`
	DeletedAt           time.Time `gorm:"index"`
}
