package app

import "github.com/fajarcandraaa/shortened_url/internal/entity"

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&entity.Url{},
	}

	return migrationData
}
