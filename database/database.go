package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("webserver.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
	}

	// migrate db
	if err := db.AutoMigrate(
		&User{},
		&Memo{},
	); err != nil {
		log.Fatal("cannot migrate DB \n", err.Error())
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	Database = DBInstance{DB: db}
}
