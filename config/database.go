package config

import (
	"os"

	"github.com/lordofthemind/ginGormCrudAPI/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	// DBHost is the host of the database.
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "keshavpostgres"
// 	dbname   = "gocrud"
// )

func DatabaseConnection() *gorm.DB {

	connectionString := os.Getenv("ELEPHANT_CONNECTION")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
