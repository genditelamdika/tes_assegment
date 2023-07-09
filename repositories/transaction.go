package repositories

import (
	"indocattes/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	Generatecsv() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	GetUserByid(ID int) (models.User, error)

	GetProductId(id int) (models.Product, error)
	Delete(id int) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Product").Find(&transactions).Error

	return transactions, err
}
func (r *repository) Generatecsv() ([]models.Transaction, error) {
	var csv []models.Transaction
	err := r.db.Preload("User").Preload("Product").Find(&csv).Error

	return csv, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Product").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Model(&transaction).Updates(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}

func (r *repository) GetProductId(id int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error

	return product, err
}
func (r *repository) GetUserByid(id int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Cart.Product").Preload("Transaction").First(&user, id).Error

	return user, err
}
func (r *repository) Delete(id int) error {
	err := r.db.Exec("DELETE FROM carts WHERE user_id = ?", id).Error

	return err
}
