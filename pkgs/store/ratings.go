package store

import (
	"errors"

	"github.com/harbur/golang-gin-starter/pkgs/models"
	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ListRatings lists ratings of a Rating
func ListRatings() []*models.Rating {
	var ratings []*models.Rating
	db.Find(&ratings)
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
