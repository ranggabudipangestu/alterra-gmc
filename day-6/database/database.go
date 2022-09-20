package database

import (
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	// we use sync.Once for make sure we create connection only once
	once sync.Once
)

// CreateConnection is a function for creating new connection with database
func CreateConnection() {

	conf := dbConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() { mysql.Connect() })
}

// GetConnection is a faction for return connection or return value dbConn
// because we set var dbConn is private
func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
