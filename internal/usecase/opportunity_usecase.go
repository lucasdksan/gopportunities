package usecase

import (
	"gopportunities/internal/repository"
	"gopportunities/internal/schemas"
)

type OpportunityUsecase struct {
	repository repository.OpportunityRepository
}

func NewOpportunityUsecase(repository repository.OpportunityRepository) OpportunityUsecase {
	return OpportunityUsecase{repository: repository}
}

func (u *OpportunityUsecase) CreateOpportunity(opening *schemas.Opening) error {
	if err := u.repository.CreateOpportunity(opening); err != nil {
		return err
	}

	return nil
}

func (u *OpportunityUsecase) DeleteOpportunity(id string, opening *schemas.Opening) error {
	if err := u.repository.DeleteOpportunity(id, opening); err != nil {
		return err
	}

	return nil
}
