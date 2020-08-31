package impl

import (
	"github.com/Scowluga/Notera/server/models"
	"github.com/Scowluga/Notera/server/persistence"
	"github.com/Scowluga/Notera/server/repositories/interfaces"
	"github.com/jinzhu/gorm"
)

// NoteRepositoryImpl ...
type NoteRepositoryImpl struct {
	db *gorm.DB
}

// NewNoteRepository generates a new note repository
func NewNoteRepository(db *gorm.DB) interfaces.NoteRepository {
	return &NoteRepositoryImpl{
		db: db,
	}
}

// CreateNote ...
func (repo *NoteRepositoryImpl) CreateNote(tx *gorm.DB, note *models.Note) error {
	return persistence.CreateNote(tx, note)
}

// GetNotesForMedia returns all notes associated with a mediaID
func (repo *NoteRepositoryImpl) GetNotesForMedia(tx *gorm.DB, mediaID string) ([]*models.Note, error) {
	return persistence.GetNotesForMedia(tx, mediaID)
}

// GetNotesForUser returns all notes associated with a userID
func (repo *NoteRepositoryImpl) GetNotesForUser(tx *gorm.DB, userID string) ([]*models.Note, error) {
	return persistence.GetNotesForUser(tx, userID)
}

// UpdateNote ...
func (repo *NoteRepositoryImpl) UpdateNote(tx *gorm.DB, note *models.Note) error {
	return persistence.UpdateNote(tx, note)
}

// DeleteNote deletes a note by ID
func (repo *NoteRepositoryImpl) DeleteNote(tx *gorm.DB, noteID string) error {
	return persistence.DeleteNote(tx, noteID)
}
