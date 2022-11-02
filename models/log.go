package models

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Token string `json:"Token" gorm:"not null"`

	// Supported Log Levels are:
	// DEBUG,INFO,WARNING,ERROR,CRITICAL,FATAL
	Level     string  `json:"Level"`
	Timestamp float64 `json:"Timestamp"`
	Log       string  `json:"Log"`
	Project   string  `json:"Project"`
	App       string  `json:"App"`
	Client    string  `json:"Client"`
}

func (l *Log) IsValid() (bool, error) {
	return true, nil
}
