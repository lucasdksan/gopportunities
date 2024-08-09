package router

import (
	"gopportunities/internal/controller"
	"gopportunities/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Initialize(opportunityUsecase usecase.OpportunityUsecase) {
	router := gin.Default()
	OpportunityController := controller.NewOpportunityController(opportunityUsecase)

	initializeRoutes(router, OpportunityController)

	router.Run(":8080")
}
