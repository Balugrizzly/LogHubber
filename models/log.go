package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	AccessToken string `json:"AccessToken" gorm:"not null"`

	// Supported Log Levels are:
	// DEBUG,INFO,WARNING,ERROR,CRITICAL,FATAL
	Level   string    `json:"Level"`
	Time    time.Time `json:"DateTime"`
	Log     string    `json:"Log"`
	Project string    `json:"Project"`
	App     string    `json:"App"`
	Client  string    `json:"Client"`
}
