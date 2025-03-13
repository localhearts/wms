package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	StatusProcessing = iota + 1
	StatusShipping
	StatusDelivered
	StatusCanceled
)

const (
	OrderTypeB2B = iota + 1
	OrderTypeB2C
)

type DeliveryOrder struct {
	DeliveryOrderID string                `gorm:"primaryKey;type:char(36);not null;" json:"delivery_order_id"`
	OrderNumber     string                `gorm:"size:50;not null;unique;" json:"order_number"`
	OrderType       int                   `gorm:"default:0;not null;" json:"order_type"`
	OrderDate       time.Time             `gorm:"not null" json:"order_date"`
	ShippingDate    *time.Time            `gorm:"default:null" json:"shipping_date"`
	CustomerID      string                `gorm:"type:char(36);not null" json:"customer_id"`
	Status          int                   `gorm:"default:0;not null;" json:"status"`
	Notes           string                `gorm:"type:text" json:"notes"`
	Details         []DeliveryOrderDetail `gorm:"foreignKey:DeliveryOrderID;" json:"details"`
	Customer        Customer              `gorm:"refrences:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"customer"`
	CreatedAt       time.Time             `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time             `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate ensures UUID is generated if it's empty
func (d *DeliveryOrder) BeforeCreate(tx *gorm.DB) (err error) {
	if d.DeliveryOrderID == "" {
		d.DeliveryOrderID = uuid.New().String()
	}
	return
}

func (d *DeliveryOrder) GetStatus() string {
	switch d.Status {
	case StatusProcessing:
		return "Processing"
	case StatusShipping:
		return "Shipping"
	case StatusDelivered:
		return "Delivered"
	case StatusCanceled:
		return "Canceled"
	default:
		return "Unknown"
	}
}

func (d *DeliveryOrder) GetOrderType() string {
	switch d.OrderType {
	case OrderTypeB2B:
		return "B2B"
	case OrderTypeB2C:
		return "B2C"
	default:
		return "Unknown"
	}
}
