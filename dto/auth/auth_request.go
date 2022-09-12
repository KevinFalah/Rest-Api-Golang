package authdto

type RegisterRequest struct {
	FullName string `json:"fullname" gorm:"type:varchar(255)" validate="required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate="required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate="required"`
	Gender   string `gorm:"type: varchar(255)" json:"gender" validate="required"`
	Phone    string `gorm:"type: varchar(255)" json:"phone" validate="required"`
	Address  string `gorm:"type: varchar(255)" json:"address" validate="required"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate="required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate="required"`
}
