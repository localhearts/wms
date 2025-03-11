package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	SupplierID    string         `gorm:"type:char(36);primaryKey;" json:"supplier_id"`
	SupplierCode  string         `gorm:"type:varchar(20);unique" json:"supplier_code"`
	SupplierName  string         `gorm:"type:varchar(255);not null" json:"supplier_name"`
	ContactPerson string         `gorm:"type:varchar(255)" json:"contact_person"`
	Phone         string         `gorm:"type:varchar(20)" json:"phone"`
	Email         string         `gorm:"type:varchar(255);unique" json:"email"` // Mencegah duplikasi email
	Address       string         `gorm:"type:varchar(255)" json:"address"`
	CityID        string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"city_id"`
	ProvinceID    string         `gorm:"type:char(36);not null;index;constraint:OnDelete:CASCADE;" json:"province_id"`
	PostalCode    string         `gorm:"type:varchar(10)" json:"postal_code"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft delete
	Products      []Product      `gorm:"foreignKey:SupplierID" json:"products"`
	City          City           `gorm:"referances:CityID" json:"city"`
	Province      Province       `gorm:"referances:ProvinceID" json:"province"`
}

// BeforeCreate untuk memastikan UUID digenerate jika belum ada
func (s *Supplier) BeforeCreate(tx *gorm.DB) (err error) {
	if s.SupplierID == "" {
		s.SupplierID = uuid.New().String()
	}
	return
}

// BeforeUpdate untuk mengupdate UpdatedAt
func (s *Supplier) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return
}

// BeforeDelete untuk menghapus relasi
func (s *Supplier) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Model(&s).Association("Products").Clear()
	return
}
