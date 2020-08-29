package api

import (
	"github.com/Scowluga/Notera/server/api/notes"
	repository_impl "github.com/Scowluga/Notera/server/repositories/impl"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// SetupServer sets up router to handle all endpoints
func SetupServer(router *mux.Router, db *gorm.DB) {
	// Initialize repos
	noteRepository := repository_impl.NewNoteRepository(db)

	// Setup endpoints
	notes.Setup(router, db, noteRepository)
}
