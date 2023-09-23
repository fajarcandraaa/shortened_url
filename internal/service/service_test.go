package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func testConfig(t *testing.T) (*gorm.DB, error) {
	err := godotenv.Load("../../.env") // Update the path accordingly
	require.NoError(t, err)
	DBDriver := os.Getenv("DB_DRIVER_TEST")
	DBHost := os.Getenv("DB_HOST_TEST")
	DBUser := os.Getenv("DB_USER_TEST")
	DBPassword := os.Getenv("DB_PASSWORD_TEST")
	DBName := os.Getenv("DB_NAME_TEST")
	DBPort := os.Getenv("DB_PORT_TEST")
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	db, err := gorm.Open(DBDriver, DBURL) // initiate database for testing
	require.NoError(t, err)
	db.AutoMigrate(&entity.Url{})

	return db, nil

}
