package models

type Transaction struct {
	ID         int    `json:"id"`
	Counterqty int    `json:"counterqty" form:"counterqty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status" gorm:"type: varchar(255)"`
	Attachment string `json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	TripID     int    `json:"tripid" form:"tripid"`
	Trip       Trip   `json:"trip" `
}
