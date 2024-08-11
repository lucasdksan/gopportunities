package configs

import (
	"gopportunities/internal/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("Postgres")

	dsn := Bank_connection_string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errf("Postgres opening error: %v", err)
		return nil, err
	}

	if err = db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Errf("Postgres automigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
