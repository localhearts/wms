package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/localhearts/wms/controllers"
	"github.com/localhearts/wms/repository"
)

func RegisterInboundRoutes(r *gin.Engine, repo repository.InboundRepository) {
	ctrl := controllers.InboundController{InboundRepo: repo}

	inbound := r.Group(os.Getenv("BASE_URL") + "/inbound")
	{
		inbound.POST("/", ctrl.CreateInbound)
		inbound.GET("/", ctrl.GetAllInbounds)
		inbound.GET("/:id", ctrl.GetInbound)
		inbound.PUT("/:id", ctrl.UpdateInbound)
		inbound.DELETE("/:id", ctrl.DeleteInbound)
	}
}
