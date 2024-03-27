package model

import "gorm.io/gorm"

type Band struct {
	gorm.Model
	Name string `gorm:"unique"`
}
