package main

import (
	"net/http"
	"os"
	"time"

	routes "github.com/Scowluga/Notera/server/api"
	"github.com/Scowluga/Notera/server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Setup DB
	db, err := setupDatabase()
	if err != nil {
		log.Warn(err)
	}
	defer db.Close()

	// Setup Server
	router := mux.NewRouter()
	err = routes.NewServer(router, db)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Starting server on port: %s", os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}

func setupDatabase() (*gorm.DB, error) {
	connStr := os.Getenv("DATABASE_CONNECTION_STRING")

	log.Infof("Connecting to PostgreSQL database: %s", os.Getenv("DBNAME"))
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	// Create Note table if not already exists
	if !db.HasTable(&models.Note{}) {
		db.CreateTable(&models.Note{})
		db.Model(&models.Note{}).Create(&models.Note{
			MediaID:   "id_Eroica",
			UserID:    "id_David",
			Timestamp: 3012020,
			Text:      "This slaaps",
		})
	}

	return db, nil
}
