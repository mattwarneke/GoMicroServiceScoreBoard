package main

import "time"

// Score an object for the user
type Score struct {
	GameScore int       `json:"score"`
	User      string    `json:"user"`
	Date      time.Time `json:"date"`
	Level     string    `json:"level"`
	Version   string    `json:"version"`
}
