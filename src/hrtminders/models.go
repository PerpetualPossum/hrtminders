package main

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint
	DiscordID   string
	Username    string
	DisplayName string
	Score       int
	LastReset   sql.NullTime
	Schedules   []Schedule
}

type Schedule struct {
	gorm.Model
	ID         uint
	Name       string
	DiscordID  string
	Cron       string
	LastRun    sql.NullTime
	NextRun    sql.NullTime
	ResetAt    sql.NullTime
	Location   sql.NullString
	Medication sql.NullString
	Dosage     sql.NullString
}
