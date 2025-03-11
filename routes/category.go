package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/localhearts/wms/controllers"
	"github.com/localhearts/wms/repository"
)

func CategoryRoutes(router *gin.Engine, repo repository.CategoryRepository) {
	ctrl := controllers.NewCategoryController(repo)
	uomGroup := router.Group(os.Getenv("BASE_URL") + "/categories")
	{
		uomGroup.GET("/", ctrl.GetDataTablesCat)
	}
}
