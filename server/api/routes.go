package api

import (
	"github.com/Scowluga/Notera/server/api/notes"
	repository_impl "github.com/Scowluga/Notera/server/repositories/impl"
	
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// SetupServer sets up router to handle all endpoints
func SetupServer(router *mux.Router, db *gorm.DB, redis *redis.Client) {
	// Initialize repos
	noteRepository := repository_impl.NewNoteRepository(db, redis)

	// Setup endpoints
	notes.Setup(router, db, noteRepository)
}
