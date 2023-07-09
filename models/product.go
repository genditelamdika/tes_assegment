package models

type Product struct {
	ID           int           `json:"id"`
	Name         string        `json:"name" form:"name" gorm:"varchar(255)"`
	Price        float64       `json:"price" form:"price"`
	Description  string        `json:"description" form:"description" gorm:"varchar(255)"`
	Image        string        `json:"image"`
	CategoryID   int           `json:"categoryid" form:"categoryid"`
	Category     Category      `json:"category"`
	Transactions []Transaction `gorm:"many2many:product_transactions;"`
	// Discountcode       string        `json:"discountcode" form:"discountcode" gorm:"varchar(255)"`
	// Discountpercentage int           `json:"discountpercentage" form:"discountpercentage"`
	// Discountamount     float64       `json:"discountamount" form:"discountamount"`
}
