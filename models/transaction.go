package models

import "time"

type Transaction struct {
	ID        int            `json:"id"`
	StartDate time.Time      `json:"-"`
	DueDate   time.Time      `json:"-"`
	Attache   string         `json:"attache"`
	Status    string         `json:"status"  gorm:"type:varchar(25)"`
	// UserID    int            `json:"userid"`
	// User      UserIdResponse `json:"user"`
}
