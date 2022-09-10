package models

import "time"

// User model struct
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Phone     int       `json:"phone" gorm:"type: int"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: text"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Subscribe bool      `json:"subscribe"`
}

// type UserIdResponse struct {
// 	ID        int    `json:"id"`
// 	FullName  string `json:"fullname"`
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// 	Gender    string `json:"gender"`
// 	Address   string `json:"address"`
// 	Phone     int    `json:"phone"`
// 	Subscribe bool   `json:"subscribe"`
// }

// func (UserIdResponse) TableName() string {
// 	return "users"
// }
