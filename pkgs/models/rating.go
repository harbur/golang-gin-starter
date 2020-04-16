package models

import "github.com/jinzhu/gorm"

// Rating rating
type Rating struct {
	gorm.Model
	Rating int   `json:"rating" binding:"required"`
	Movie  Movie `json:"movie" binding:"required" gorm:"association_autoupdate:false"`
}
