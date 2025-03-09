package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	StorageStatusAvailable = iota + 1
	StorageStatusFull
	StorageStatusMaintenance
)

type Storage struct {
	StorageID   string         `gorm:"type:char(36);primaryKey;" json:"storage_id"`
	WarehouseID string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"warehouse_id"`
	Location    string         `gorm:"not null" json:"location"`
	Rack        string         `gorm:"not null" json:"rack"`
	Level       int            `gorm:"not null" json:"level"`
	Bin         int            `gorm:"not null" json:"bin"`
	Capacity    int            `gorm:"not null" json:"capacity"`
	CurrentLoad int            `gorm:"not null" json:"current_load"`
	Temperature float64        `json:"temperature,omitempty"`
	Status      int            `gorm:"default:0" json:"status"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft delete
}

// BeforeCreate untuk memastikan UUID digenerate jika belum ada
func (s *Storage) BeforeCreate(tx *gorm.DB) (err error) {
	if s.StorageID == "" {
		s.StorageID = uuid.New().String()
	}
	return
}

// GetStorageStatusString is a function to get storage status in string
func (s *Storage) GetStorageStatusString() string {
	switch s.Status {
	case StorageStatusAvailable:
		return "Available"
	case StorageStatusFull:
		return "Full"
	case StorageStatusMaintenance:
		return "Maintenance"
	default:
		return "Unknown"
	}
}
