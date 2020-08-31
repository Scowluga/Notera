package notes

import (
	"net/http"

	repository_interfaces "github.com/Scowluga/Notera/server/repositories/interfaces"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type NoteSubRouter struct {
	Router         *mux.Router
	Db             *gorm.DB
	noteRepository repository_interfaces.NoteRepository
}

// Setup for notes endpoint
func Setup(router *mux.Router, db *gorm.DB, noteRepository repository_interfaces.NoteRepository) {
	note := NoteSubRouter{
		noteRepository: noteRepository,
	}

	note.Router = router

	note.Db = db

	note.Router.HandleFunc("/notes/media/{mediaID}", note.MediaHandler).Methods("GET")
}

// MediaHandler handles getting all notes for a specific media
func (sr *NoteSubRouter) MediaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedMediaID := vars["mediaID"]

	tx := sr.Db.Begin()
	notes, err := sr.noteRepository.GetNotesForMedia(tx, requestedMediaID)
	if err != nil {
		tx.Rollback()

	}

	// w.WriteHeader(http.StatusOK)
	// m := map[string]string{"Boo": "Yes Babe"}
	// json.NewEncoder(w).Encode(m)
}
