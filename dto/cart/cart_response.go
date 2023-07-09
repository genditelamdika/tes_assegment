package cartsdto

import "indocattes/models"

type CartResponse struct {
	ID        int            `json:"id"`
	ProductID int            `json:"productid" `
	Product   models.Product `json:"product" form:"product" gorm:"foreignKey:ProductID"`
	UserID    int            `json:"userid"  `
	User      models.User    `json:"user" gorm:"foreignKey:UserID"`
	Status    string         `json:"status" gorm:"type: varchar(255)"`
}
