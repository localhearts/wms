package repository

import (
	// Ganti dengan path package models Anda

	"github.com/localhearts/wms/models"
	"gorm.io/gorm"
)

// UomRepository mendefinisikan interface repository untuk Uom.
type UomRepository interface {
	// GetDataTablesUom mengambil data Uom dengan parameter pagination, filtering, dan sorting.
	GetDataTablesUom(start int, length int, searchValue string, orderBy string) (uoms []models.Uom, totalRecords int64, filteredRecords int64, err error)
}

type uomRepository struct {
	DB *gorm.DB
}

// NewUomRepository mengembalikan instance baru dari UomRepository.
func NewUomRepository(db *gorm.DB) UomRepository {
	return &uomRepository{DB: db}
}

// GetDataTablesUom mengimplementasikan query DataTables.
func (r *uomRepository) GetDataTablesUom(start int, length int, searchValue string, orderBy string) (uoms []models.Uom, totalRecords int64, filteredRecords int64, err error) {
	// Hitung total data Uom tanpa filter.
	if err = r.DB.Model(&models.Uom{}).Count(&totalRecords).Error; err != nil {
		return
	}

	// Siapkan query untuk filter.
	query := r.DB.Model(&models.Uom{})
	if searchValue != "" {
		query = query.Where("uom_name LIKE ?", "%"+searchValue+"%")
	}

	// Hitung data setelah filter.
	if err = query.Count(&filteredRecords).Error; err != nil {
		return
	}

	// Ambil data dengan pengurutan, offset, dan limit.
	err = query.Order(orderBy).Offset(start).Limit(length).Find(&uoms).Error
	return
}
