package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Enum Status untuk Purchase Order
const (
	StatusPending = iota + 1
	StatusApproved
	StatusRejected
)

type PurchaseOrder struct {
	PurchaseOrderID string                `gorm:"type:char(36);primaryKey;" json:"purchase_order_id"`
	PurchaseOrderNo string                `gorm:"size:20;not null;uniqueIndex" json:"purchase_order_no"` // Contoh: PO-20250308-001
	SupplierID      string                `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"supplier_id"`
	OrderDate       time.Time             `gorm:"not null" json:"order_date"`
	ExpectedDate    *time.Time            `json:"expected_date,omitempty"` // Bisa NULL jika tidak ditentukan
	Status          int                   `gorm:"default:0" json:"status"` // 0: Pending, 1: Approved, 2: Rejected
	TotalAmount     float64               `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	CreatedBy       string                `gorm:"size:50;not null" json:"created_by"`
	ApprovedBy      *string               `gorm:"size:50" json:"approved_by,omitempty"` // Nullable
	Remarks         *string               `gorm:"type:text" json:"remarks,omitempty"`   // Nullable
	Details         []PurchaseOrderDetail `gorm:"foreignKey:PurchaseOrderID;" json:"purchase_details"`

	// Relasi dengan Supplier
	Supplier  Supplier       `gorm:"foreignKey:SupplierID;" json:"supplier"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Hook untuk generate UUID sebelum Create
func (po *PurchaseOrder) BeforeCreate(tx *gorm.DB) (err error) {
	if po.PurchaseOrderID == "" {
		po.PurchaseOrderID = uuid.New().String()
	}
	return
}

// Method untuk mengubah Status ke String agar lebih mudah dibaca di API
func (po *PurchaseOrder) GetStatus() string {
	switch po.Status {
	case StatusPending:
		return "Pending"
	case StatusApproved:
		return "Approved"
	case StatusRejected:
		return "Rejected"
	default:
		return "Unknown"
	}
}
