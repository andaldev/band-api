package db

import (
	"band_api/helper"
	"band_api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	helper.ErrorPanic(err)

	// database.Migrator().DropTable(&model.Band{})
	// database.Clauses(clause.OnConflict{DoNothing: true})
	database.AutoMigrate(&model.Band{})

	// loop this with a slice
	database.FirstOrCreate(&model.Band{Name: "Opeth"})
	database.FirstOrCreate(&model.Band{Name: "Katatonia"})
	database.FirstOrCreate(&model.Band{Name: "Tool"})
	database.FirstOrCreate(&model.Band{Name: "Deftones"})

	return database
}
