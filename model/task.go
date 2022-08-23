package model

import "time"

type Task struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:true" json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rate        string    `json:"rate"`
	DateID      uint64    `gorm:"not null" json:"-"`
	Date        Date      `json:"foreignkey:DateID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"date"`
	StateID     uint64    `gorm:"not null" json:"-"`
	State       State     `json:"foreignkey:StateID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"state"`
}

type Date struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Year      int    `json:"year"`
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	StartHour int    `json:"starthour"`
	EndHour   int    `json:"endhour"`
}

type State struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Canceled  bool   `json:"canceled"`
	Finished  bool   `json:"finished"`
	Continues bool   `json:"continues"`
}

type DateNumber struct {
	Day   []int
	Month []int
	Year  []int
	Hour  []int
}

type CurrentTime struct {
	Day   int
	Month int
	Year  int
	Hour  int
}

type TaskDTO struct {
	Title       string
	Description string
	Date        Date
	Rate        string
	// State       State
}
