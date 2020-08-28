package models

import "github.com/jinzhu/gorm"

// Note represents an annotation on a Spotify track with timestamp
type Note struct {
	gorm.Model
	SpotifyTrackID string `gorm:"type:character varying"`
	Text           string `gorm:"type:character varying"`
}
