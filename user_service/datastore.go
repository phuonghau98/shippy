package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host = "localhost"
	user = "phuong"
	DBName = "shippy"
	password = "1"
	port = "27019"
)
func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	//host := os.Getenv("DB_HOST")
	//user := os.Getenv("DB_USER")
	//DBName := os.Getenv("DB_NAME")
	//password := os.Getenv("DB_PASSWORD")
	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			host, port, user, DBName, password,
		),
	)
}