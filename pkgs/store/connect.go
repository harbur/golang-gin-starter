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
	// Please define your user name and password for my sql.
	d, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db = d
	db.AutoMigrate(&models.Movie{})
}
