package model

import "gorm.io/gorm"

type Band struct {
	gorm.Model
	Id   uint
	Name string
}
