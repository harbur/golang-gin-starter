package models

// Movie movie
type Movie struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" binding:"required"`
}

// ListMovies REST response listing movies
type ListMovies struct {
	Movies []*Movie `json:"movies"`
}

// GetMovie REST response getting movie
type GetMovie struct {
	Movie Movie `json:"movie"`
}
