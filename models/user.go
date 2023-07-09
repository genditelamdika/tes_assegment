package models

import "time"

type User struct {
	ID           int           `json:"id"`
	Fullname     string        `json:"fullname" gorm:"type: varchar(255)"`
	Email        string        `json:"email" gorm:"type: varchar(255)"`
	Password     string        `json:"-" gorm:"type: varchar(255)"`
	Role         string        `json:"role" gorm:"type: varchar(255)"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	Cart         []Cart        `json:"cart"`
	Transactions []Transaction `gorm:"many2many:user_transactions;"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
