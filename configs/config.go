package configs

import (
	"gopportunities/internal/tools"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *tools.Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *tools.Logger {
	logger = tools.NewLogger(prefix)
	return logger
}
