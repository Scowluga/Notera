package util

import (
	"github.com/Scowluga/Notera/server/models"
)

type NoteResponse struct {
	ID 			uint 
	MediaID 	string 
	UserID 		string 
	Timestamp 	int 
	Text 		string
}

func GenerateNoteResponse(note *models.Note) NoteResponse {
	return NoteResponse {
		ID: 		note.ID,
		MediaID: 	note.MediaID,
		UserID: 	note.UserID, 
		Timestamp: 	note.Timestamp,
		Text: 		note.Text,
	}
}
