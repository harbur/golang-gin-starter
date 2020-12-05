package store

import (
	"log"

	"github.com/harbur/golang-gin-starter/pkgs/models"
	"github.com/jinzhu/gorm"

	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// MoviesDB movies database
type MoviesDB interface {
	ListMovies() models.ListMovies
	CreateMovie(movie models.Movie) (models.Movie, error)
	GetMovie(id uint) (models.GetMovie, error)
	UpdateMovie(id uint, movie models.Movie) (models.Movie, error)
	DeleteMovie(id uint)
}

type moviesDB struct {
	db *gorm.DB
}

// Movies DB instance
var Movies MoviesDB

func init() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	db.AutoMigrate(&models.Movie{})
	Movies = &moviesDB{db: db}
}
