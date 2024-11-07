package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDatabase(DBMigrator func(db *gorm.DB) error) {
	dsn := os.Getenv("dbstring")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		os.Exit(2)
	}
	log.Println("Connected to database")

	err = DBMigrator(db)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		os.Exit(2)
	}

	Database = DbInstance{Db: db}
}
