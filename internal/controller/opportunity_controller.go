package controller

import (
	"gopportunities/configs"
	"gopportunities/internal/model"
	"gopportunities/internal/schemas"
	"gopportunities/internal/tools"
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
	request := model.CreateOpeningRequest{}
	logger := configs.GetLogger("Main")

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Validation error %v", err.Error())
		tools.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if errCreatingOpportunity := c.OpportunityUsecase.CreateOpportunity(&opening); errCreatingOpportunity != nil {
		logger.Errf("error creating opening: %v", errCreatingOpportunity.Error())
		tools.SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	tools.SendSuccess(ctx, "create-opening", opening)
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
