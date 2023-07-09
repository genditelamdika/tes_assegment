package models

import "time"

type Transaction struct {
	ID                 int       `json:"id"`
	Status             string    `json:"status" form:"status" gorm:"type: varchar(255)"`
	Date               time.Time `json:"-"`
	UserID             int       `json:"userid"`
	User               User      `json:"user"`
	Discountcode       string    `json:"discountcode" form:"discountcode" gorm:"varchar(255)"`
	Discountpercentage int       `json:"discountpercentage" form:"discountpercentage"`
	Discountamount     float64   `json:"discountamount" form:"discountamount"`
	Product            []Product `json:"product"  gorm:"many2many:product_transactions;"`
	Total              int       `json:"total"`
	Qty                int       `json:"qty"`
}
