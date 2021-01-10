package impl

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/Scowluga/Notera/server/models"
	"github.com/Scowluga/Notera/server/persistence"
	"github.com/Scowluga/Notera/server/repositories/interfaces"
	
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

const redisExpiryTime = 24
const redisNoteKey = "note"

// NoteRepositoryImpl ...
type NoteRepositoryImpl struct {
	db 		*gorm.DB
	redis 	*redis.Client
}

// NewNoteRepository generates a new note repository
func NewNoteRepository(db *gorm.DB, redis *redis.Client) interfaces.NoteRepository {
	return &NoteRepositoryImpl{
		db: db,
		redis: redis,
	}
}

// CreateNote ...
func (repo *NoteRepositoryImpl) CreateNote(tx *gorm.DB, note *models.Note) error {
	repo.redis.Del(redisNoteKey)
	return persistence.CreateNote(tx, note)
}

// GetNotesForMedia returns all notes associated with a mediaID
func (repo *NoteRepositoryImpl) GetNotesForMedia(tx *gorm.DB, mediaID string) ([]*models.Note, error) {
	encodedQuery := base64.StdEncoding.EncodeToString([]byte(mediaID))
	notesJson, err := repo.redis.HGet(redisNoteKey, encodedQuery).Result()

	if err == nil && len(notesJson) > 0 {
		log.Infof("Fetching mediaID from redis: " + mediaID)

		var notes []*models.Note 
		err =  json.Unmarshal([]byte(notesJson), &notes)
		if err != nil {
			log.Warn(err)
			return nil, err
		}

		return notes, nil
	} else {
		log.Infof("Cache miss, querying DB")
		notes, err := persistence.GetNotesForMedia(tx, mediaID)

		// Attempt to cache result
		notesJson, err := json.Marshal(notes)
		if err != nil {
			log.Warn("Error marshalling result to json: " + err.Error())
			return notes, nil
		}

		log.Infof("Caching result")
		repo.redis.HSet(redisNoteKey, encodedQuery, notesJson)
		repo.redis.Expire(redisNoteKey, redisExpiryTime * time.Hour)

		return notes, nil
	}
}

// GetNotesForUser returns all notes associated with a userID
func (repo *NoteRepositoryImpl) GetNotesForUser(tx *gorm.DB, userID string) ([]*models.Note, error) {
	encodedQuery := base64.StdEncoding.EncodeToString([]byte(userID))
	notesJson, err := repo.redis.HGet(redisNoteKey, encodedQuery).Result()

	if err == nil && len(notesJson) > 0 {
		log.Infof("Fetching userID from redis: " + userID)

		var notes []*models.Note 
		err =  json.Unmarshal([]byte(notesJson), &notes)
		if err != nil {
			log.Warn(err)
			return nil, err
		}

		return notes, nil
	} else {
		log.Infof("Cache miss, querying DB")
		notes, err := persistence.GetNotesForUser(tx, userID)

		// Attempt to cache result
		notesJson, err := json.Marshal(notes)
		if err != nil {
			log.Warn("Error marshalling result to json: " + err.Error())
			return notes, nil
		}

		log.Infof("Caching result")
		repo.redis.HSet(redisNoteKey, encodedQuery, notesJson)
		repo.redis.Expire(redisNoteKey, redisExpiryTime * time.Hour)

		return notes, nil
	}
}

// UpdateNote ...
func (repo *NoteRepositoryImpl) UpdateNote(tx *gorm.DB, note *models.Note) error {
	repo.redis.Del(redisNoteKey)
	return persistence.UpdateNote(tx, note)
}

// DeleteNote deletes a note by ID
func (repo *NoteRepositoryImpl) DeleteNote(tx *gorm.DB, noteID string) error {
	repo.redis.Del(redisNoteKey)
	return persistence.DeleteNote(tx, noteID)
}
