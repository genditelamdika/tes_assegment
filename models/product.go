package models

type Product struct {
	ID          int     `json:"id" gorm:"primary_key:auto_increment"`
	Name        string  `json:"name" form:"name" gorm:"varchar(255)"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description" gorm:"varchar(255)"`
	Image       string  `json:"image"`
	// CategoryID []int      `json:"category_id" form:"category_id" gorm:"-"`
	// Category   []Category `json:"category" gorm:"many2many:product_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
