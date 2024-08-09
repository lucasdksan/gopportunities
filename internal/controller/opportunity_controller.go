package controller

import (
	"gopportunities/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OpportunityController struct {
	OpportunityUsecase usecase.OpportunityUsecase
}

func NewOpportunityController(usecase usecase.OpportunityUsecase) OpportunityController {
	return OpportunityController{
		OpportunityUsecase: usecase,
	}
}

func (c *OpportunityController) GetOpportunity(ctx *gin.Context)    {}
func (c *OpportunityController) CreateOpportunity(ctx *gin.Context) {}
func (c *OpportunityController) UpdateOpportunity(ctx *gin.Context) {}
func (c *OpportunityController) DeleteOpportunity(ctx *gin.Context) {}
func (c *OpportunityController) ListOpportunity(ctx *gin.Context)   {}
