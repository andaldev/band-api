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

	database.AutoMigrate(&model.Band{})

	// loop this with a slice
	database.Create(&model.Band{Name: "Opeth"})
	database.Create(&model.Band{Name: "Katatonia"})
	database.Create(&model.Band{Name: "Tool"})
	database.Create(&model.Band{Name: "Deftones"})

	// var band model.Band
	// database.First(&band, "name = ?", "Opeth")

	return database
}
