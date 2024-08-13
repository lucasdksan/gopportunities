package repository

import (
	"gopportunities/internal/model"
	"gopportunities/internal/schemas"

	"gorm.io/gorm"
)

type OpportunityRepository struct {
	connection *gorm.DB
}

func NewOpportunityRepository(connnection *gorm.DB) OpportunityRepository {
	return OpportunityRepository{
		connection: connnection,
	}
}

func (r *OpportunityRepository) CreateOpportunity(opening *schemas.Opening) error {
	if err := r.connection.Create(opening).Error; err != nil {
		return err
	}

	return nil
}

func (r *OpportunityRepository) DeleteOpportunity(id string, opening *schemas.Opening) error {
	if err := r.connection.Delete(opening, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *OpportunityRepository) ListOpportunity() ([]model.Opening, error) {
	openings := []model.Opening{}

	if err := r.connection.Find(&openings).Error; err != nil {
		return []model.Opening{}, err
	}

	return openings, nil
}

func (r *OpportunityRepository) GetOpportunity(id string) (model.Opening, error) {
	opening := model.Opening{}

	if err := r.connection.First(&opening, id).Error; err != nil {
		return model.Opening{}, err
	}

	return opening, nil
}
