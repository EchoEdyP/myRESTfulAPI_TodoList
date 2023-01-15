// Package database contains the interface and implementation for connecting to the database
package database

import (
	"database/sql"
	"github.com/kelseyhightower/envconfig"
)

// DBConn is an interface for connecting to the database
type DBConn interface {
	Connect() (*sql.DB, error)
}

// MySQLConn is an implementation of DBConn that connects to a MySQL database
type MySQLConn struct{}

type Config struct {
	DBDriver string `envconfig:"DB_DRIVER" default:"mysql"`
	DBUser   string `envconfig:"DB_USER" default:"eep"`
	DBPass   string `envconfig:"DB_PASS" default:"1903"`
	DBName   string `envconfig:"DB_NAME" default:"RESTfulAPI_todos"`
}

// Connect connects to a MySQL database using the environment variables specified in .env
func (c *MySQLConn) Connect() (*sql.DB, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(config.DBDriver, config.DBUser+":"+config.DBPass+"@/"+config.DBName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
