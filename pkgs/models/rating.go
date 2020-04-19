package models

import "github.com/jinzhu/gorm"

// Rating rating
type Rating struct {
	gorm.Model
	Rating  int    `json:"rating" binding:"required"`
	MovieID int    `json:"-"`
	Movie   *Movie `json:"-" gorm:"association_autoupdate:false"`
}
