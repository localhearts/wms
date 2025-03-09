package repository

import (
	"github.com/google/uuid"
	"github.com/localhearts/wms/models"
	"gorm.io/gorm"
)

type InboundRepository struct {
	DB *gorm.DB
}

// Create
func (repo *InboundRepository) CreateInbound(inbound *models.Inbound) error {
	return repo.DB.Create(inbound).Error
}

// Get by ID
func (repo *InboundRepository) GetInboundByID(id uuid.UUID) (*models.Inbound, error) {
	var inbound models.Inbound
	err := repo.DB.Preload("Customer").First(&inbound, "inbound_id = ?", id).Error
	return &inbound, err
}

// Get All
func (repo *InboundRepository) GetAllInbounds() ([]models.Inbound, error) {
	var inbounds []models.Inbound
	err := repo.DB.Preload("Customer").Find(&inbounds).Error
	return inbounds, err
}

// Update
func (repo *InboundRepository) UpdateInbound(inbound *models.Inbound) error {
	return repo.DB.Model(&models.Inbound{}).
		Where("inbound_id = ?", inbound.InboundID).
		Updates(inbound).Error
}

// Delete (soft delete)
func (repo *InboundRepository) DeleteInbound(id uuid.UUID) error {
	return repo.DB.Delete(&models.Inbound{}, "inbound_id = ?", id).Error
}
