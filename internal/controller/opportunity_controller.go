package controller

import (
	"gopportunities/internal/usecase"
	"net/http"

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

func (c *OpportunityController) GetOpportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}

func (c *OpportunityController) CreateOpportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}

func (c *OpportunityController) UpdateOpportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

func (c *OpportunityController) DeleteOpportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}

func (c *OpportunityController) ListOpportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "LIST Opening",
	})
}
