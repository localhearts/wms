package models

import "time"

type InboundDetail struct {
	InboundDetailID  uint `gorm:"primaryKey"`
	InboundID        uint
	ProductID        uint
	Quantity         int
	ReceivedQuantity int       `gorm:"default:0"`
	Status           string    `gorm:"type:enum('Pending','Received','Placed');default:'Pending'"`
	inbound          Inbound   `gorm:"foreignKey:InboundID"`
	product          Product   `gorm:"foreignKey:ProductID"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
	DeletedAt        time.Time `gorm:"index"`
}
