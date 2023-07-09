package transactionsdto

type TransactionResponse struct {
	ID                 int     `json:"id"`
	Date               string  `json:"date" form:"date" gorm:"varchar(255)" validate:"required"`
	Status             string  `json:"status" form:"status" gorm:"varchar(255)" validate:"required"`
	UserID             int     `json:"userid"`
	Discountcode       string  `json:"discountcode" form:"discountcode" gorm:"varchar(255)"`
	Discountpercentage int     `json:"discountpercentage" form:"discountpercentage"`
	Discountamount     float64 `json:"discountamount" form:"discountamount"`
	// ProductID int    `json:"productid"`
	Total int    `json:"total"`
	Qty   string `json:"qty"`
}
