package main

import (
	"gopportunities/configs"
	"gopportunities/internal/database"
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

	dbConnection, err := database.ConnectDB()

	if errConfig := configs.Init(); errConfig != nil {
		logger.Errf("Config initialization error: %v", err)
		return
	}

	if err != nil {
		panic(err)
	}

	OpportunityRepository := repository.NewOpportunityRepository(dbConnection)
	OpportunityUseCase := usecase.NewOpportunityUsecase(OpportunityRepository)

	router.Initialize(OpportunityUseCase)
}
