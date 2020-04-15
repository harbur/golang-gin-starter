package store

import (
	"github.com/harbur/golang-gin-starter/pkgs/models"
	"github.com/jinzhu/gorm"

	// sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

// Connect connects to db
func Connect() {
	d, _ := gorm.Open("sqlite3", ":memory:")
	db = d
	db.AutoMigrate(&models.Movie{})
}
