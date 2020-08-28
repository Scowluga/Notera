package main

import (
	"encoding/json"
	"net/http"

	"github.com/Scowluga/Notera/server/db"
	"github.com/Scowluga/Notera/server/models"

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
	db, err := db.SetupDatabase()
	if err != nil {
		log.Warn(err)
	}
	defer db.Close()

	// API
	router := mux.NewRouter()
	router.HandleFunc("/boo", MyHandler)

	// ---
	var note models.Note
	db.Take(&note)

	log.Infof("Note: %s %s \n", note.SpotifyTrackID, note.Text)

}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	m := map[string]string{"Boo": "Yes Babe"}
	json.NewEncoder(w).Encode(m)
}
