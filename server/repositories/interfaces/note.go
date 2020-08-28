package interfaces

import (
	"github.com/Scowluga/Notera/server/models"
	"github.com/jinzhu/gorm"
)

// NoteRepository ...
type NoteRepository interface {
	CreateNote(tx *gorm.DB, note *models.Note) error
	// TODO: add pagination
	GetNotesForMedia(tx *gorm.DB, mediaID string) ([]*models.Note, error)
	GetNotesForUser(tx *gorm.DB, userID string) ([]*models.Note, error)
	UpdateNote(tx *gorm.DB, note *models.Note) error
	DeleteNote(tx *gorm.DB, noteID string) error
}
