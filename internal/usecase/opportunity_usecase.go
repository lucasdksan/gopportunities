package usecase

import (
	"gopportunities/internal/model"
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

func (u *OpportunityUsecase) ListOpportunity() ([]model.Opening, error) {
	openings, err := u.repository.ListOpportunity()

	if err != nil {
		return []model.Opening{}, err
	}

	return openings, nil
}

func (u *OpportunityUsecase) GetOpportunity(id string) (model.Opening, error) {
	opening, err := u.repository.GetOpportunity(id)

	if err != nil {
		return model.Opening{}, err
	}

	return opening, nil
}

func (u *OpportunityUsecase) UpdateOpportunity(opening *schemas.Opening) error {
	if err := u.repository.UpdateOpportunity(opening); err != nil {
		return err
	}

	return nil
}

func (u *OpportunityUsecase) ValidateOpportunity(id string, opening *schemas.Opening) error {
	if err := u.repository.ValidateOpportunity(id, opening); err != nil {
		return err
	}

	return nil
}
