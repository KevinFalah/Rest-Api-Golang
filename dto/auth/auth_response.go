package authdto

type RegisterResponse struct {
	FullName  string `json:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"password" gorm:"type: varchar(255)"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe bool   `json:"subscribe" gorm:"type: varchar(255)"`
}

type LoginResponse struct {
	ID       int    `json:"id" gorm:"type: int"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
