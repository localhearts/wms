package repository

import (
	// Ganti dengan path package models Anda

	"github.com/localhearts/wms/models"
	"gorm.io/gorm"
)

// CategoryRepository mendefinisikan interface repository untuk Categories.
type CategoryRepository interface {
	// GetDataTablesUom mengambil data Uom dengan parameter pagination, filtering, dan sorting.
	GetDataTablesCat(start int, length int, searchValue string, orderBy string) (uoms []models.Category, totalRecords int64, filteredRecords int64, err error)
}

type catRepository struct {
	DB *gorm.DB
}

// NewUomRepository mengembalikan instance baru dari UomRepository.
func NewCatRepository(db *gorm.DB) CategoryRepository {
	return &catRepository{DB: db}
}

// GetDataTablesUom mengimplementasikan query DataTables.
func (r *catRepository) GetDataTablesCat(start int, length int, searchValue string, orderBy string) (uoms []models.Category, totalRecords int64, filteredRecords int64, err error) {
	// Hitung total data Uom tanpa filter.
	if err = r.DB.Model(&models.Category{}).Count(&totalRecords).Error; err != nil {
		return
	}

	// Siapkan query untuk filter.
	query := r.DB.Model(&models.Category{})
	if searchValue != "" {
		query = query.Where("category_name LIKE ?", "%"+searchValue+"%")
	}

	// Hitung data setelah filter.
	if err = query.Count(&filteredRecords).Error; err != nil {
		return
	}

	// Ambil data dengan pengurutan, offset, dan limit.
	err = query.Order(orderBy).Offset(start).Limit(length).Find(&uoms).Error
	return
}
