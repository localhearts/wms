package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/localhearts/wms/models"
	"github.com/localhearts/wms/repository"
)

type InboundController struct {
	InboundRepo repository.InboundRepository
}

// Create Inbound
func (ctrl *InboundController) CreateInbound(c *gin.Context) {
	var inbound models.Inbound
	if err := c.ShouldBindJSON(&inbound); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.InboundBeforeCreate(&inbound); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.InboundRepo.CreateInbound(&inbound)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, inbound)
}

// Get Inbound by ID
func (ctrl *InboundController) GetInbound(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}
	inbound, err := ctrl.InboundRepo.GetInboundByID(uuidID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Inbound not found"})
		return
	}
	c.JSON(200, inbound)
}

// Get All Inbounds
func (ctrl *InboundController) GetAllInbounds(c *gin.Context) {
	inbounds, err := ctrl.InboundRepo.GetAllInbounds()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, inbounds)
}

// Update Inbound
func (ctrl *InboundController) UpdateInbound(c *gin.Context) {
	var inbound models.Inbound
	if err := c.ShouldBindJSON(&inbound); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.InboundRepo.UpdateInbound(&inbound); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, inbound)
}

// Delete Inbound
func (ctrl *InboundController) DeleteInbound(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}
	if err := ctrl.InboundRepo.DeleteInbound(uuidID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}

// InboundBeforeCreate automatically generates a UUID for inbound if not provided.
func (ctrl *InboundController) InboundBeforeCreate(inbound *models.Inbound) error {
	if inbound.InboundID == "" {
		inbound.InboundID = uuid.New().String()
	}
	inbound.CreatedAt = time.Now()
	return nil
}
