// Package database contains the interface and implementation for connecting to the database
package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	"os"
)

// DBConn is an interface for connecting to the database
type DBConn interface {
	Connect() (*sql.DB, error)
}

// MySQLConn is an implementation of DBConn that connects to a MySQL database
type MySQLConn struct{}

// Connect connects to a MySQL database using the environment variables specified in .env
func (c *MySQLConn) Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
