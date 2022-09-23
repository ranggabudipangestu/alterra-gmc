package database

import (
	"sync"

	"github.com/ranggabudipangestu/hexagonal-api/pkg/util"
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
		User: util.Getenv("DB_USER", "root"),
		Pass: util.Getenv("DB_PASS", "root"),
		Host: util.Getenv("DB_HOST", "localhost"),
		Port: util.Getenv("DB_PORT", "3306"),
		Name: util.Getenv("DB_NAME", "db-test"),
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
