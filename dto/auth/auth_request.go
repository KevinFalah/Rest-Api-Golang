package authdto

//	type AuthRequest struct {
//		ID       int    `json:"id"`
//		Name     string `gorm:"type: varchar(255)" json:"name"`
//		Email    string `gorm:"type: varchar(255)" json:"email"`
//		Password string `gorm:"type: varchar(255)" json:"password"`
//	}

type RegisterRequest struct {
	FullName string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"addres" gorm:"type: varchar(255)"`
}
