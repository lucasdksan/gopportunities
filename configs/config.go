package configs

import (
	"fmt"
	"gopportunities/internal/tools"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *tools.Logger
)

func Init() error {
	var err error

	InitializeEnv()

	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("error initializing postegres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(prefix string) *tools.Logger {
	logger = tools.NewLogger(prefix)
	return logger
}
