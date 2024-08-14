package test

import (
	"os"
	"testing"

	"gopportunities/configs"

	"github.com/stretchr/testify/assert"
)

func TestInitializeEnv(t *testing.T) {
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_USER", "user")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_DB", "test_db")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("API_PORT", "8080")

	configs.InitializeEnv()

	assert.Equal(t, "localhost", configs.Host_db)
	assert.Equal(t, "user", configs.User_db)
	assert.Equal(t, "password", configs.Password_db)
	assert.Equal(t, "test_db", configs.Database)
	assert.Equal(t, 5432, configs.Port_db)
	assert.Equal(t, 8080, configs.Port)
}
