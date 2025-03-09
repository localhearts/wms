package models

type OutboundDetail struct {
	DetailID       uint `gorm:"primaryKey"`
	OutboundID     uint
	ProductID      uint
	Quantity       int
	PickedQuantity int      `gorm:"default:0"`
	Status         string   `gorm:"type:enum('Pending','Picked','Packed');default:'Pending'"`
	Outbound       Outbound `gorm:"foreignKey:OutboundID"`
	Product        Product  `gorm:"foreignKey:ProductID"`
}
