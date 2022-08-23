package model

import "time"

// ----------> Model created for the user.
type User struct {
	ID        uint      `json:"userID" gorm:"primary_key:auto_increment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required" gorm:"type:varchar(255); uniqueIndex"`
	Password  string    `json:"pass"  gorm:"->;<-; not null"` // allow read and write
	// Tasks     []Task    `json:"tasks"`
}

// ----------> ModelDTO created for the user.
type UserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"type:varchar(255); uniqueIndex"`
	Password string `json:"-"  gorm:"->;<-; not null"`
}
