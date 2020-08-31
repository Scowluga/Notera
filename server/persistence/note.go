package persistence

import (
	"github.com/Scowluga/Notera/server/models"
	"github.com/jinzhu/gorm"
)

// CreateNote ...
func CreateNote(tx *gorm.DB, note *models.Note) error {
	return tx.Create(note).Error
}

// GetNotesForMedia returns all notes associated with a mediaID
func GetNotesForMedia(tx *gorm.DB, mediaID string) ([]*models.Note, error) {
	var notes []*models.Note

	err := tx.Where("media_id = ?", mediaID).Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

// GetNotesForUser returns all notes associated with a userID
func GetNotesForUser(tx *gorm.DB, userID string) ([]*models.Note, error) {
	var notes []*models.Note

	err := tx.Where("user_id = ?", userID).Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

// UpdateNote ...
func UpdateNote(tx *gorm.DB, note *models.Note) error {
	return tx.Save(note).Error
}

// DeleteNote deletes a note by ID
func DeleteNote(tx *gorm.DB, noteID string) error {
	return tx.Where("id = ?", noteID).Delete(&models.Note{}).Error
}
