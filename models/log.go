package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	AccessToken string `json:"AccessToken" gorm:"not null"`

	// Supported Log Levels are:
	// DEBUG,INFO,WARNING,ERROR,CRITICAL,FATAL
	Level     string `json:"Level"`
	Timestamp uint   `json:"Timestamp"`
	Log       string `json:"Log"`
	Project   string `json:"Project"`
	App       string `json:"App"`
	Client    string `json:"Client"`
}

func (l *Log) IsValid() (bool, error) {
	return true, nil
}
