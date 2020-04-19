package store

import (
	"errors"

	"github.com/harbur/golang-gin-starter/pkgs/models"

	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ListRatings lists ratings of a Rating
func ListRatings(id uint) []*models.Rating {
	var ratings []*models.Rating
	// movie := models.Movie{Model: gorm.Model{ID: id}}
	db.Model(id).Find(&ratings)
	return ratings
}

// CreateRating creates a rating
func CreateRating(rating models.Rating) (models.Rating, error) {
	// make sure payload does not contain id
	if rating.ID != 0 {
		return rating, errors.New("invalid id")
	}
	err := db.Create(&rating).Error
	return rating, err
}

// GetRating gets a rating
func GetRating(id uint) (models.Rating, error) {
	var rating models.Rating
	err := db.Where("ID = ?", id).Find(&rating).Error
	return rating, err
}

// UpdateRating updates a rating
func UpdateRating(id uint, rating models.Rating) (models.Rating, error) {
	// make sure payload contains correct id
	if id != rating.ID {
		return rating, errors.New("invalid id")
	}
	// make sure rating exists
	_, err := GetRating(id)
	if err != nil {
		return rating, err
	}
	err = db.Save(&rating).Error
	return rating, err
}

// DeleteRating deletes a rating
func DeleteRating(id uint) {
	var rating models.Rating
	db.Where("ID = ?", id).Delete(rating)
}
