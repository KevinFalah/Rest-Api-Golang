package transactionsdto

// import "time"

type TransactionResponse struct {
	ID        int    `json:"id"`
	StartDate string `json:"startdate"`
	DueDate   string `json:"duedate"`
	UserID    int    `json:"user_id" form:"user_id"`
	Attache   string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    bool   `json:"status" gorm:"type:text" form:"status"`
}
