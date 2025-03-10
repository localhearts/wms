package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/localhearts/wms/models"
	"github.com/localhearts/wms/repository"
)

type UomController struct {
	repo repository.UomRepository
}

func NewUomController(repo repository.UomRepository) *UomController {
	return &UomController{repo}
}

// CreateUom menangani pembuatan data Uom baru
func (ctrl *UomController) CreateUom(c *gin.Context) {
	var uom models.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.repo.Create(&uom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, uom)
}

// GetUom menampilkan data Uom berdasarkan ID
func (ctrl *UomController) GetUom(c *gin.Context) {
	id := c.Param("id")
	uom, err := ctrl.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "UOM tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, uom)
}

// GetAllUom menampilkan semua data Uom
func (ctrl *UomController) GetAllUom(c *gin.Context) {
	uoms, err := ctrl.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, uoms)
}

// UpdateUom mengubah data Uom yang sudah ada
func (ctrl *UomController) UpdateUom(c *gin.Context) {
	id := c.Param("id")
	existing, err := ctrl.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "UOM tidak ditemukan"})
		return
	}

	// Bind data request ke object existing
	if err := c.ShouldBindJSON(existing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Pastikan ID tidak berubah
	existing.UomID = id

	if err := ctrl.repo.Update(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, existing)
}

// DeleteUom menghapus data Uom berdasarkan ID
func (ctrl *UomController) DeleteUom(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "UOM berhasil dihapus"})
}
