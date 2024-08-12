package controller

import (
	"fmt"
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
	id := ctx.Query("id")

	if id == "" {
		tools.SendError(ctx, http.StatusBadRequest, tools.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := c.OpportunityUsecase.DeleteOpportunity(id, &opening); err != nil {
		tools.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	tools.SendSuccess(ctx, "delete-opening", opening)
}

func (c *OpportunityController) ListOpportunity(ctx *gin.Context) {
	openings, err := c.OpportunityUsecase.ListOpportunity()

	if err != nil {
		tools.SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	tools.SendSuccess(ctx, "listing-openigs", openings)
}
