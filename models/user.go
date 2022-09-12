package models

import "time"

type User struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"addres" gorm:"type: varchar(255)"`
	Subscribe bool      `json:"subscribe" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Subscribe bool      `json:"subscribe"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UserAuthResponse struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (UserResponse) TableName() string {
	return "users"
}

func (UserAuthResponse) TableName() string {
	return "users"
}
