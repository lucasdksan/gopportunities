package main

import (
	"gopportunities/internal/database"
	"gopportunities/internal/repository"
	"gopportunities/internal/router"
	"gopportunities/internal/usecase"
)

func main() {
	dbConnection, err := database.ConnectDB()

	if err != nil {
		panic(err)
	}

	OpportunityRepository := repository.NewOpportunityRepository(dbConnection)
	OpportunityUseCase := usecase.NewOpportunityUsecase(OpportunityRepository)

	router.Initialize(OpportunityUseCase)
}
