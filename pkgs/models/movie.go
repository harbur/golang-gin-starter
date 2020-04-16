package models

import "github.com/jinzhu/gorm"

// Movie movie
type Movie struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}
