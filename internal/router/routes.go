package router

import (
	"gopportunities/internal/controller"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine, handler controller.OpportunityController) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpportunity)
		v1.POST("/opening", handler.CreateOpportunity)
		v1.DELETE("/opening", handler.DeleteOpportunity)
		v1.PUT("/opening", handler.UpdateOpportunity)
		v1.GET("/opening/list", handler.ListOpportunity)
	}
}
