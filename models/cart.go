package models

type Cart struct {
	ID            int     `json:"id`
	ProductID     int     `json:"productid`
	Product       Product `json:"product" form:"product" gorm:"foreignKey:ProductID`
	UserID        int     `json:"userid`
	User          User    `json:"user" gorm:"foreignKey:UserID`
	Status string  `json:"status" gorm:"type: varchar(255)"`
}
