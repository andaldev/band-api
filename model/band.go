package model

import "gorm.io/gorm"

type Band struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// type Band struct {
// 	ID   int    `gorm:"unique;primaryKey;autoIncrement"`
// 	Name string `gorm:"type:varchar(255)"`
// }
