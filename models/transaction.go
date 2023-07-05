package models

type Transaction struct {
	ID     int    `json:"name" form:"name" gorm:"varchar(255)"`
	Date   string `json:"date" form:"date" gorm:"varchar(255)"`
	Status string `json:"status" form:"status" gorm:"varchar(255)"`
}
