package models

import "time"

type DeliveryOrderDetail struct {
	DeliveryOrderDetailID string        `gorm:"type:char(36);primaryKey;not null;" json:"delivery_order_detail_id"`
	DeliveryOrderID       string        `gorm:"type:char(36);index;not null;constraint:OnDelete:CASCADE;" json:"delivery_order_id"`
	ProductID             string        `gorm:"type:char(36);index;not null;constraint:OnDelete:CASCADE;" json:"product_id"`
	DeliveryOrder         DeliveryOrder `gorm:"references:DeliveryOrderID;" json:"delivery_order"`
	Product               Product       `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	UomID                 string        `gorm:"type:char(36);index;not null;constraint:OnDelete:CASCADE;" json:"uom_id"`
	Uom                   Uom           `gorm:"refrences:UomID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"uom"`
	Quantity              int           `gorm:"not null" json:"quantity"`
	Notes                 string        `gorm:"type:text" json:"notes"`
	CreatedAt             time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}
