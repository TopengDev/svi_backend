package models

import "gorm.io/gorm"


type Post struct {
	gorm.Model
	Title string `gorm:"not null;size:200;"`
	Content string `gorm:"not null;"`
	Category string `gorm:"not null;size:100;"`
	Created_date string
	Updated_date string
	Status string `gorm:"not null;size:100;"`

}
