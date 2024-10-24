package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	Title  string
	Author string
}
