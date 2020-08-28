package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// NewServer sets up router to handle all endpoints
func NewServer(router *mux.Router, db *gorm.DB) error {
	router.HandleFunc("/boo", BooHandler)
	return nil
}

// BooHandler returns the single JSON result "Boo": "Yes Babe"
func BooHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	m := map[string]string{"Boo": "Yes Babe"}
	json.NewEncoder(w).Encode(m)
}
