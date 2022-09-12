package transactionsdto

type TransactionRequest struct {
	StartDate string `json:"startdate" gorm:"type: varchar(255)"`
	DueDate   string `json:"duedate" gorm:"type: varchar(255)"`
	UserID    int    `json:"user_id" form:"user_id"`
	Attache   string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool   `json:"status"`
}

type CreateTransactionRequest struct {
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	Attache   string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool   `json:"status" form:"status"`
}

type UpdateTransactionRequest struct {
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	Attache   string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool   `json:"status"`
}
