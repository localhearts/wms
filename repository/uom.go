package repository

import (
	"errors"

	"github.com/localhearts/wms/models"

	"gorm.io/gorm"
)

type UomRepository interface {
	Create(uom *models.Uom) error
	GetByID(id string) (*models.Uom, error)
	GetAll() ([]models.Uom, error)
	Update(uom *models.Uom) error
	Delete(id string) error
}

type uomRepository struct {
	db *gorm.DB
}

func NewUomRepository(db *gorm.DB) UomRepository {
	return &uomRepository{db}
}

func (r *uomRepository) Create(uom *models.Uom) error {
	return r.db.Create(uom).Error
}

func (r *uomRepository) GetByID(id string) (*models.Uom, error) {
	var uom models.Uom
	result := r.db.First(&uom, "uom_id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &uom, result.Error
}

func (r *uomRepository) GetAll() ([]models.Uom, error) {
	var uoms []models.Uom
	result := r.db.Find(&uoms)
	return uoms, result.Error
}

func (r *uomRepository) Update(uom *models.Uom) error {
	return r.db.Save(uom).Error
}

func (r *uomRepository) Delete(id string) error {
	return r.db.Delete(&models.Uom{}, "uom_id = ?", id).Error
}
