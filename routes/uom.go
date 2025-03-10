package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/localhearts/wms/controllers"
	"github.com/localhearts/wms/repository"
)

func UomRoutes(router *gin.Engine, repo repository.UomRepository) {
	ctrl := controllers.NewUomController(repo)
	uomGroup := router.Group(os.Getenv("BASE_URL") + "/uom")
	{
		uomGroup.GET("/", ctrl.GetDataTablesUom)
	}
}
