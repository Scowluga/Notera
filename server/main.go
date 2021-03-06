package main

import (
	"net/http"
	"os"
	"fmt"

	routes "github.com/Scowluga/Notera/server/api"
	"github.com/Scowluga/Notera/server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/go-redis/redis/v7"
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
		log.Fatal(err)
	}
	defer db.Close()

	// Setup Redis
	redis := setupCache()

	// Setup Server
	router := mux.NewRouter()
	if router == nil {
		log.Fatal("Router creation returns null")
	}
	routes.SetupServer(router, db, redis)

	port := os.Getenv("SERVER_PORT")
	log.Infof("Starting server on port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func setupDatabase() (*gorm.DB, error) {
	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
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
			Timestamp: 420,
			Text:      "This slaaps",
		})
		log.Infof("Connected to empty postgres, created notes table and pre-populated with single note")
	} else {
		log.Infof("Connected to postgres")
	}

	return db, nil
}

func setupCache() *redis.Client {
	client := redis.NewClient(&redis.Options {
		Addr: 		os.Getenv("REDIS_URL"),
		Password:	"",
		DB: 		0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Warn("Couldn't establish redis client with error: " + err.Error())
	} else {
		log.Infof("Connected to redis: " + pong)
	}

	return client
}
