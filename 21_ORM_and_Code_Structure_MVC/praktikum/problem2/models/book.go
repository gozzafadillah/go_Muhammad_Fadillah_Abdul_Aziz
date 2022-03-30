package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}
