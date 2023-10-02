package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the User model.
type User struct {
	gorm.Model
	PhoneNumber   string         `gorm:"type:varchar(13);uniqueIndex;not null"`
	FullName      string         `gorm:"type:varchar(60);not null"`
	PasswordHash  string         `gorm:"type:text;not null"`
	LoginAttempts []LoginAttempt // Relationship: a user has many login attempts
}

// LoginAttempt represents the LoginAttempt model.
type LoginAttempt struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	LoginTime  time.Time
	Successful bool `gorm:"not null"`
}

type Result struct {
	ID           int    `json:"id"`
	PasswordHash string `json:"passwordHash"`
	FullName     string `json:"fullName"`
	PhoneNumber  string `json:"phoneNumber"`
}
