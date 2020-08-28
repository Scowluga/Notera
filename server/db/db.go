package db

import (
	"os"

	"github.com/Scowluga/Notera/server/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // To connect using Postgres
	log "github.com/sirupsen/logrus"
)

// SetupDatabase connects to and sets up the Postgres database
func SetupDatabase() (*gorm.DB, error) {
	connStr := os.Getenv("DATABASE_CONNECTION_STRING")

	log.Infof("Connecting to PostgreSQL database with string: %s", connStr)
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	// Create Note table if not already exists
	if !db.HasTable(&models.Note{}) {
		db.CreateTable(&models.Note{})
		db.Model(&models.Note{}).Create(&models.Note{SpotifyTrackID: "Eroica", Text: "This slaaps"})
	}

	return db, nil
}
