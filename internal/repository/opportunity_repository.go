package repository

import (
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
