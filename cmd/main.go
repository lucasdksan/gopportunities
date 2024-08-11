package main

import (
	"gopportunities/configs"
	"gopportunities/internal/repository"
	"gopportunities/internal/router"
	"gopportunities/internal/tools"
	"gopportunities/internal/usecase"
)

var (
	logger tools.Logger
)

func main() {
	logger = *configs.GetLogger("Main")

	if errConfig := configs.Init(); errConfig != nil {
		logger.Errf("Config initialization error: %v", errConfig)
		return
	}

	dbConnection := configs.GetPostgres()

	OpportunityRepository := repository.NewOpportunityRepository(dbConnection)
	OpportunityUseCase := usecase.NewOpportunityUsecase(OpportunityRepository)

	router.Initialize(OpportunityUseCase)
}
