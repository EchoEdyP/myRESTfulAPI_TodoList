package database

import (
	"database/sql"
	"fmt"
)

type Config struct {
	DBDriver string `envconfig:"DB_DRIVER" default:"mysql"`
	DBUser   string `envconfig:"DB_USER" default:"eep"`
	DBPass   string `envconfig:"DB_PASS" default:"1903"`
	DBName   string `envconfig:"DB_NAME" default:"RESTfulAPI_todos"`
}

func ConnectDB(config *Config) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@/%s",
		config.DBUser,
		config.DBPass,
		config.DBName,
	)
	db, err := sql.Open(config.DBDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}
