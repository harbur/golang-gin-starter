package store

import (
	"testing"

	"github.com/harbur/golang-gin-starter/pkgs/models"
	"github.com/stretchr/testify/assert"
)

// TestListRatingsOK lists ratings
func TestListRatingsOK(t *testing.T) {
	// Prepare
	Connect()
	movie := models.Movie{
		Name: "godfather",
	}
	movie, err := CreateMovie(movie)
	assert.NoError(t, err)

	rating := models.Rating{
		Rating: 1,
		Movie:  movie,
	}

	// movies list is empty
	ratings := ListRatings()
	assert.Len(t, ratings, 0)

	// create a rating
	rating, err = CreateRating(rating)
	assert.NoError(t, err)

	// movies list has a movie
	ratings = ListRatings()
	assert.Len(t, ratings, 1)
	// assert.Equal(t, rating.Movie.ID, ratings[0].Movie.ID)
}
