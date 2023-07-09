package repositories

import (
	"fmt"
	"indocattes/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCart() ([]models.Cart, error)
	// FindCartID(ID int) ([]models.Cart, error)
	GetCartById(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	// UpdateProduct(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)

	GetUserByID(ID int) (*models.User, error)
	GetPendingProducts(ID int) ([]string, error)
	FindcartByUser(ID int) ([]models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCart() ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("User").Find(&carts).Error

	return carts, err
}

func (r *repository) FindcartByUser(ID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("User").Find(&carts, "user_id = ?", ID).Error
	fmt.Println(ID)
	return carts, err
}

func (r *repository) FindCartID(ID int) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product").Preload("User.Cart").Find(&carts, "user_id = ?", ID).Error
	fmt.Println(ID)
	return carts, err
}

func (r *repository) GetCartById(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Preload("User").Create(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) GetUserByID(ID int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetPendingProducts(ID int) ([]string, error) {
	var products []models.Product
	err := r.db.Where("user_id = ? AND status = ?", ID, "pending").Find(&products).Error
	if err != nil {
		return nil, err
	}
	var pendingProducts []string
	for _, product := range products {
		pendingProducts = append(pendingProducts, product.Name)
	}
	return pendingProducts, nil
}
