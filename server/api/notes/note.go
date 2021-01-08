package notes

import (
	"net/http"
	
	"github.com/Scowluga/Notera/server/api/util"
	repository_interfaces "github.com/Scowluga/Notera/server/repositories/interfaces"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type NoteSubRouter struct {
	Router         *mux.Router
	Db             *gorm.DB
	noteRepository repository_interfaces.NoteRepository
}

type NoteResponse struct {
	ID 			uint 
	MediaID 	string 
	UserID 		string 
	Timestamp 	int 
	Text 		string
}

// Setup for notes endpoint
func Setup(router *mux.Router, db *gorm.DB, noteRepository repository_interfaces.NoteRepository) {
	note := NoteSubRouter{
		noteRepository: noteRepository,
	}

	note.Router = router

	note.Db = db

	note.Router.HandleFunc("/notes/media/{mediaID}", note.MediaHandler).Methods("GET")
	note.Router.HandleFunc("/notes/user/{userID}", note.UserHandler).Methods("GET")
}

// MediaHandler handles getting all notes for a specific media
func (sr *NoteSubRouter) MediaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedMediaID := vars["mediaID"]
	log.Infof("MediaHandler - requested media id: " + requestedMediaID)

	notes, err := sr.noteRepository.GetNotesForMedia(sr.Db, requestedMediaID)
	if err != nil {
		log.WithError(err).Warn("MediaHandler")
	}
	
	res := make([]*NoteResponse, 0, len(notes))
	for _, note := range notes {
		res = append(res, &NoteResponse{
			ID: 		note.ID,
			MediaID: 	note.MediaID,
			UserID: 	note.UserID, 
			Timestamp: 	note.Timestamp,
			Text: 		note.Text,
		})
	}

	util.Response(w, http.StatusOK, map[string]interface{} {
		"notes": res,
	})
}

// UserHandler handles getting all notes for a specific user
func (sr *NoteSubRouter) UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedUserID := vars["userID"]
	log.Infof("UserID - requested user id: " + requestedUserID)

	notes, err := sr.noteRepository.GetNotesForUser(sr.Db, requestedUserID)
	if err != nil {
		log.WithError(err).Warn("UserHandler")
	}
	
	res := make([]*NoteResponse, 0, len(notes))
	for _, note := range notes {
		res = append(res, &NoteResponse{
			ID: 		note.ID,
			MediaID: 	note.MediaID,
			UserID: 	note.UserID, 
			Timestamp: 	note.Timestamp,
			Text: 		note.Text,
		})
	}

	util.Response(w, http.StatusOK, map[string]interface{} {
		"notes": res,
	})
}
