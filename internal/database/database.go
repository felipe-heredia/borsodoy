package database

import (
	"borsodoy/radovid/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var config = &gorm.Config{TranslateError: true}

func Initialize() {
	var err error
	Database, err = gorm.Open(sqlite.Open("internal/database/database.db"), config)

	if err != nil {
		log.Fatal("failed to connect database")
	}

	Database.AutoMigrate(new(models.User), new(models.Item))
}

func SetupTestDB() {
	var err error
	Database, err = gorm.Open(sqlite.Open(":memory:"), config)

	if err != nil {
		log.Fatal("failed to connect database")
	}

	Database.AutoMigrate(new(models.User), new(models.Item))
}
