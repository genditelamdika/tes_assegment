package models

type Order struct {
	ID           int    `json:"id"  gorm:"primary_key:auto_increment"`
	Date         string `json:"date"`
	Status       string `json:"status" gorm:"varchar(255)"`
	DiscountCode string `json:"discount_code"`
	// CustomerID   int              `json:"customer_id"`
	// Customer     CustomerResponse `json:"customer"`
	// ProductID    int              `json:"product_id"`
	// Product      ProductResponse  `json:"product"`
}
