package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string  `gorm:"uniqueIndex;not null"`
	Email        string  `gorm:"uniqueIndex;not null"`
	Password     string  `gorm:"not null"`
	GoogleID     *string `gorm:"uniqueIndex"`
	LastLoginAt  time.Time
	TodoItems    []TodoItem
	Priorities   []Priority
	Contacts     []Contact
	WaterIntakes []WaterIntake
	Thoughts     []Thought
}

type TodoItem struct {
	gorm.Model
	UserID      uint
	Title       string `gorm:"not null"`
	Description string
	DueDate     time.Time
	Completed   bool `gorm:"default:false"`
}

type Priority struct {
	gorm.Model
	UserID      uint
	Title       string `gorm:"not null"`
	Description string
	Date        time.Time
	Completed   bool `gorm:"default:false"`
}

type Contact struct {
	gorm.Model
	UserID      uint
	Name        string `gorm:"not null"`
	Type        string `gorm:"not null"` // Call, Email, or Text
	Description string
	Date        time.Time
	Completed   bool `gorm:"default:false"`
}

type WaterIntake struct {
	gorm.Model
	UserID  uint
	Date    time.Time
	Glasses int `gorm:"default:0"`
	Target  int `gorm:"default:10"`
}

type Thought struct {
	gorm.Model
	UserID  uint
	Content string `gorm:"not null"`
	Date    time.Time
}
