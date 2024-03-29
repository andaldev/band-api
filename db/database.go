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

	initialBands := []model.Band{
		{Name: "Opeth"},
		{Name: "Katatonia"},
		{Name: "Tool"},
		{Name: "Deftones"},
		{Name: "Alice in Chains"},
	}

	for _, band := range initialBands {
		err := database.Where(&band).FirstOrCreate(&band).Error
		helper.ErrorPanic(err)
	}

	return database
}
