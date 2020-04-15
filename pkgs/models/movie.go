package models

// Movie movie
type Movie struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name" binding:"required"`
}
