package models

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
