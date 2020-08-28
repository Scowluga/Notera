package models

import "github.com/jinzhu/gorm"

// Note represents an annotation on a Spotify track with timestamp
type Note struct {
	gorm.Model
	MediaID   string `json:"mediaID" gorm:"type:TEXT"`
	UserID    string `json:"userID" gorm:"type:TEXT"`
	Timestamp int    `json:"timestamp" gorm:"type:INT"`
	Text      string `json:"text" gorm:"type:TEXT"`
}
