package models

import "time"

type Transaction struct {
	ID        int          `json:"id" gorm:"primary_key:auto_increment"`
	StartDate string       `json:"startdate" form:"startdate" gorm:"type: varchar(255)"`
	DueDate   string       `json:"duedate" form:"duedate" gorm:"type: varchar(255)"`
	User      UserResponse `json:"user"`
	UserID    int          `json:"user_id" form:"user_id"`
	Attache   string       `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool         `json:"status" gorm:"type:text" form:"status"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}

type TransactionResponse struct {
	ID        int          `json:"id"`
	StartDate string       `json:"startdate"`
	DueDate   string       `json:"duedate"`
	User      UserResponse `json:"user"`
	UserID    int          `json:"user_id"`
	Attache   string       `json:"attache"`
	Status    bool         `json:"status"`
}

type TransactionUserResponse struct {
	ID        int    `json:"id"`
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id"`
	Attache   string `json:"attache"`
	Status    bool   `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}

func (TransactionUserResponse) TableName() string {
	return "transactions"
}
