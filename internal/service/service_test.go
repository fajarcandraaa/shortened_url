package service_test

import (
	"testing"

	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func testConfig(t *testing.T) (*gorm.DB, error) {
	var (
		dsn = "host=localhost user=postgres dbname=db_shortened_url_test sslmode=disable password=postgres port=5433"
	)

	db, err := gorm.Open("postgres", dsn) // initiate database for testing
	require.NoError(t, err)
	db.AutoMigrate(&entity.Url{})

	err = godotenv.Load("../../.env") // Update the path accordingly
	if err != nil {
		return nil, err
	}

	return db, nil

}
