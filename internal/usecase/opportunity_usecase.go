package usecase

import "gopportunities/internal/repository"

type OpportunityUsecase struct {
	repository repository.OpportunityRepository
}

func NewOpportunityUsecase(repository repository.OpportunityRepository) OpportunityUsecase {
	return OpportunityUsecase{repository: repository}
}
