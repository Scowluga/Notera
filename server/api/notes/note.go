package notes

import (
	"encoding/json"
	"net/http"

	repository_interfaces "github.com/Scowluga/Notera/server/repositories/interfaces"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Setup for notes endpoint
func Setup(router *mux.Router, db *gorm.DB, noteRepository repository_interfaces.NoteRepository) {
	router.HandleFunc("/boo", BooHandler)
}

// BooHandler returns the single JSON result "Boo": "Yes Babe"
func BooHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	m := map[string]string{"Boo": "Yes Babe"}
	json.NewEncoder(w).Encode(m)
}
