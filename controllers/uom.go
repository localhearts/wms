package controllers

import (
	"net/http"
	"strconv"

	// Ganti dengan path package repository Anda

	"github.com/gin-gonic/gin"
	"github.com/localhearts/wms/repository"
)

// DataTablesResponse adalah struktur respons yang sesuai dengan DataTables dan tambahan informasi pagination.
type DataTablesResponse struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int64       `json:"recordsTotal"`
	RecordsFiltered int64       `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
	TotalPage       int         `json:"totalPage"`
	CurrentPage     int         `json:"currentPage"`
}

// UomController memiliki dependency ke repository Uom.
type UomController struct {
	repo repository.UomRepository
}

// NewUomController membuat instance baru UomController.
func NewUomController(repo repository.UomRepository) *UomController {
	return &UomController{repo: repo}
}

// GetDataTablesUom menampilkan data Uom dengan pagination, pencarian, dan pengurutan.
func (ctrl *UomController) GetDataTablesUom(c *gin.Context) {
	// Ambil parameter dari query string.
	draw, _ := strconv.Atoi(c.DefaultQuery("draw", "1"))
	start, _ := strconv.Atoi(c.DefaultQuery("start", "0"))
	length, _ := strconv.Atoi(c.DefaultQuery("length", "10"))
	searchValue := c.Query("search[value]")

	// Tentukan kolom pengurutan berdasarkan parameter DataTables.
	orderColumn := c.DefaultQuery("order[0][column]", "1")
	orderDir := c.DefaultQuery("order[0][dir]", "asc")
	var orderBy string
	// Misalnya, mapping kolom: 0 -> uom_id, 1 -> uom_name
	switch orderColumn {
	case "0":
		orderBy = "uom_id"
	case "1":
		orderBy = "uom_name"
	default:
		orderBy = "uom_id"
	}
	orderBy = orderBy + " " + orderDir

	// Panggil method repository untuk mengambil data.
	uoms, totalRecords, filteredRecords, err := ctrl.repo.GetDataTablesUom(start, length, searchValue, orderBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hitung total halaman dan halaman saat ini
	totalPage := 0
	if length > 0 {
		totalPage = int((filteredRecords + int64(length) - 1) / int64(length))
	}
	currentPage := (start / length) + 1

	// Siapkan response sesuai format DataTables dengan tambahan pagination
	response := DataTablesResponse{
		Draw:            draw,
		RecordsTotal:    totalRecords,
		RecordsFiltered: filteredRecords,
		Data:            uoms,
		TotalPage:       totalPage,
		CurrentPage:     currentPage,
	}

	c.JSON(http.StatusOK, response)
}
