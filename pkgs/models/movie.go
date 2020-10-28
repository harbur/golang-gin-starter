package models

// Movie movie
type Movie struct {
	ID   uint   `gorm:"primary_key"`
	Name string `json:"name" binding:"required"`
}
