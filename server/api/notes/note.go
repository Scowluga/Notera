package notes

import (
	"net/http"
	"encoding/json"
	
	"github.com/Scowluga/Notera/server/api/util"
	"github.com/Scowluga/Notera/server/models"
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

func Setup(router *mux.Router, db *gorm.DB, noteRepository repository_interfaces.NoteRepository) {
	note := NoteSubRouter{
		Router: router,
		Db: db,
		noteRepository: noteRepository,
	}

	note.Router.HandleFunc("/notes/", note.CreateHandler).Methods("POST")

	note.Router.HandleFunc("/notes/media/{mediaID}", note.MediaHandler).Methods("GET")
	note.Router.HandleFunc("/notes/user/{userID}", note.UserHandler).Methods("GET")

	note.Router.HandleFunc("/notes/", note.UpdateHandler).Methods("PUT")
	note.Router.HandleFunc("/notes/{noteID}", note.DeleteHandler).Methods("DELETE")
}

func (sr *NoteSubRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var note models.Note 

	if err := decoder.Decode(&note); err != nil {
		log.WithError(err).Warn("CreateHandler")
		util.Respond(w, http.StatusBadRequest, util.Message(err.Error()))
	} 
	
	if err := sr.noteRepository.CreateNote(sr.Db, &note); err != nil {
		log.WithError(err).Warn("CreateHandler")
		util.Respond(w, http.StatusInternalServerError, util.Message(err.Error()))
	}

	util.Respond(w, http.StatusOK, map[string]interface{} {
		"note": util.GenerateNoteResponse(&note),
	})
}

// MediaHandler handles getting all notes for a specific media
func (sr *NoteSubRouter) MediaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestedMediaID := vars["mediaID"]
	log.Infof("MediaHandler - requested media id: " + requestedMediaID)

	notes, err := sr.noteRepository.GetNotesForMedia(sr.Db, requestedMediaID)
	if err != nil {
		log.WithError(err).Warn("MediaHandler")
		util.Respond(w, http.StatusInternalServerError, util.Message(err.Error()))
	}
	
	res := make([]*util.NoteResponse, 0, len(notes))
	for _, note := range notes {
		noteResponse := util.GenerateNoteResponse(note)
		res = append(res, &noteResponse)
	}

	util.Respond(w, http.StatusOK, map[string]interface{} {
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
		util.Respond(w, http.StatusInternalServerError, util.Message(err.Error()))
	}
	
	res := make([]*util.NoteResponse, 0, len(notes))
	for _, note := range notes {
		noteResponse := util.GenerateNoteResponse(note)
		res = append(res, &noteResponse)
	}

	util.Respond(w, http.StatusOK, map[string]interface{} {
		"notes": res,
	})
}

func (sr *NoteSubRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var note models.Note 

	if err := decoder.Decode(&note); err != nil {
		log.WithError(err).Warn("UpdateHandler")
		util.Respond(w, http.StatusBadRequest, util.Message(err.Error()))
	} 
	
	if err := sr.noteRepository.UpdateNote(sr.Db, &note); err != nil {
		log.WithError(err).Warn("UpdateHandler")
		util.Respond(w, http.StatusInternalServerError, util.Message(err.Error()))
	}

	util.Respond(w, http.StatusOK, map[string]interface{} {
		"note": util.GenerateNoteResponse(&note),
	})
}

func (sr *NoteSubRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID := vars["noteID"]

	if err := sr.noteRepository.DeleteNote(sr.Db, noteID); err != nil {
		log.WithError(err).Warn("DeleteHandler")
		util.Respond(w, http.StatusInternalServerError, util.Message(err.Error()))
	}	

	util.Empty(w)
}