package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Config holds the database configuration
type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

func Initialize(db Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", db.User, db.Password, db.Host, db.Port, db.DBName)
	fmt.Println("Connection string:", connStr)
	database, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}
