package cartsdto

import "indocattes/models"

type CreateCartRequest struct {
	// ID        int     `json:"id"`
	ProductID int            `json:"productid" `
	Product   models.Product `json:"product" form:"product" gorm:"foreignKey:ProductID"`
	Status    string         `json:"status" gorm:"type: varchar(255)"`
}

type UpdateCartRequest struct {
	ID        int `json:"id"`
	ProductID int `json:"productid" `
	// // Product   Product `json:"product" form:"product" gorm:"foreignKey:ProductID"`
	// UserID int `json:"userid"  `
	// // User      User    `json:"user" gorm:"foreignKey:UserID"`
}
