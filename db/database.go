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

	database.Migrator().DropTable(&model.Band{})
	database.AutoMigrate(&model.Band{})

	// loop this with a slice
	database.Where(&model.Band{Name: "Opeth"}).FirstOrCreate(&model.Band{Name: "Opeth"})
	database.Where(&model.Band{Name: "Katatonia"}).FirstOrCreate(&model.Band{Name: "Katatonia"})
	database.Where(&model.Band{Name: "Tool"}).FirstOrCreate(&model.Band{Name: "Tool"})
	database.Where(&model.Band{Name: "Deftones"}).FirstOrCreate(&model.Band{Name: "Deftones"})

	return database
}
