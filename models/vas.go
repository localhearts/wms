package models

import (
	"time"
)

type Vas struct {
	VASID       uint `gorm:"primaryKey"`
	InboundID   uint
	ServiceType string    `gorm:"type:enum('Kitting','Merger','Labeling','Packaging')"`
	ProcessTime time.Time `gorm:"autoCreateTime"`
	PIC         string    `gorm:"type:varchar(255)"`
	Status      string    `gorm:"type:enum('Pending','Completed');default:'Pending'"`
	Inbound     Inbound   `gorm:"foreignKey:InboundID"`
}
