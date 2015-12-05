package model

import (
	"time"
)

type Todo struct {
	ID          uint       `gorm:"primary_key" sql:"not null" json:"id"`
	UserID      uint       `sql:"not null;index" json:"-"`
	Title       string     `sql:"not null" json:"title" type:"varchar(50)" binding:"required,max=50"`
	Completed   bool       `sql:"not null" json:"completed"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `sql:"not null" json:"created_at"`
}

func (u *Todo) BeforeSave() {
	if !u.Completed {
		u.CompletedAt = nil
	}
}
