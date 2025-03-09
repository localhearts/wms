package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Enum Status untuk Purchase Order Detail
const (
	StatusWaiting = iota + 1
	StatusReadyToSend
	StatusSoldOut
)

type PurchaseOrderDetail struct {
	PurchaseOrderDetailID string `gorm:"type:char(36);primaryKey;" json:"purchase_order_detail_id"`
	PurchaseOrderID       string `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"purchase_order_id"`
	ProductID             string `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"product_id"`
	Quantity              int    `gorm:"not null" json:"quantity"`
	Status                int    `gorm:"default:0" json:"status"` // 0: Waiting, 1: Ready To Send, 2: Sold Out

	// Relasi dengan Purchase Order & Product
	Product   Product        `gorm:"foreignKey:ProductID" json:"product"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Hook untuk Generate UUID Otomatis
func (pod *PurchaseOrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	if pod.PurchaseOrderDetailID == "" {
		pod.PurchaseOrderDetailID = uuid.New().String()
	}
	return nil
}

// Method untuk Mengembalikan Status dalam Bentuk String
func (pod *PurchaseOrderDetail) GetStatus() string {
	switch pod.Status {
	case StatusWaiting:
		return "Waiting"
	case StatusReadyToSend:
		return "Ready To Send"
	case StatusSoldOut:
		return "Sold Out"
	default:
		return "Unknown"
	}
}
