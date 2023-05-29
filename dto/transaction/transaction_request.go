package transactiondto

import "tour/models"

type CreateTransactionRequest struct {
	Counterqty int         `json:"counterqty" form:"counterqty"`
	Total      int         `json:"total" form:"total"`
	Status     string      `json:"status" form:"status" gorm:"type: varchar(255)"`
	Attachment string      `json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	TripID     int         `json:"tripid" form:"tripid"`
	Trip       models.Trip `json:"trip" form:"form" `
}

type UpdateTransactionRequest struct {
	Counterqty int         `json:"counterqty" form:"counterqty"`
	Total      int         `json:"total" form:"total"`
	Status     string      `json:"status" form:"status" gorm:"type: varchar(255)"`
	Attachment string      `json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	TripID     int         `json:"tripid" form:"tripid"`
	Trip       models.Trip `json:"trip" form:"trip" validate:"required"`
	// Category    models.Category `json:"category" form:"category" validate:"required"
}
