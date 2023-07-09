package transactionsdto

type CreateTransactionRequest struct {
	// Date   string `json:"date" form:"date" gorm:"varchar(255)"`
	Status             string  `json:"status" form:"status" gorm:"varchar(255)"`
	UserID             int     `json:"userid"`
	Discountcode       string  `json:"discountcode" form:"discountcode" gorm:"varchar(255)"`
	Discountpercentage int     `json:"discountpercentage" form:"discountpercentage"`
	Discountamount     float64 `json:"discountamount" form:"discountamount"`
	// ProductID int    `json:"productid"`
	Total int    `json:"total"`
	Qty   string `json:"qty"`
}

type UpdateTransactionRequest struct {
	Date   string `json:"date" form:"date" gorm:"varchar(255)"`
	Status string `json:"status" form:"status" gorm:"varchar(255)"`
	UserID int    `json:"userid"`
	Total  string `json:"total"`
	Qty    string `json:"qty"`
}
