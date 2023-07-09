package productsdto

import "indocattes/models"

type CreateProductRequest struct {
	Name        string          `json:"name" form:"name" gorm:"varchar(255)"`
	Price       float64         `json:"price" form:"price"`
	Description string          `json:"description" form:"description" gorm:"varchar(255)"`
	Image       string          `json:"image"`
	CategoryID  int             `json:"categoryid" form:"categoryid" gorm:"-"`
	Category    models.Category `json:"category" gorm:"many2many:product_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Discountcode       string          `json:"discountcode" form:"discountcode" gorm:"varchar(255)"`
	// Discountpercentage int             `json:"discountpercentage" form:"discountpercentage"`
	// Discountamount     float64         `json:"discountamount" form:"discountamount"`
	// Transaction []Transaction `json:"transaction"`
}

type UpdateProductRequest struct {
	// ID          int           `json:"id" gorm:"primary_key:auto_increment"`
	Name        string          `json:"name" form:"name" gorm:"varchar(255)"`
	Price       float64         `json:"price" form:"price"`
	Description string          `json:"description" form:"description" gorm:"varchar(255)"`
	Image       string          `json:"image"`
	CategoryID  int             `json:"categoryid" form:"categoryid" gorm:"-"`
	Category    models.Category `json:"category" gorm:"many2many:product_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Transaction []Transaction   `json:"transaction"`
}
