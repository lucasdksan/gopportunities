package repository

import "database/sql"

type OpportunityRepository struct {
	connection *sql.DB
}

func NewOpportunityRepository(connnection *sql.DB) OpportunityRepository {
	return OpportunityRepository{
		connection: connnection,
	}
}
