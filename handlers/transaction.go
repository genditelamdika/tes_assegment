package handlers

import (
	"encoding/csv"
	"fmt"
	dto "indocattes/dto/result"
	transactionsdto "indocattes/dto/transaction"
	"indocattes/models"
	"indocattes/repositories"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Discount struct {
	Code     string
	Discount float64
	Category string
	Days     []time.Weekday
}
type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}
func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}
func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// responseCategory := categoriesdto.CategoryResponse{
	// 	ID:   transaction.ID,
	// 	Name: transaction.Name,
	// 	Film: []categoriesdto.CategoryFilm{},
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

var discounts = []Discount{
	{
		Code:     "IC003",
		Discount: 0.1,
	},
	{
		Code:     "IC042",
		Discount: 0.05,
		Category: "elektronik",
	},
	{
		Code:     "IC015",
		Discount: 0.1,
		Days:     []time.Weekday{time.Saturday, time.Sunday},
	},
}

func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	request := new(transactionsdto.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.TransactionRepository.GetUserByid(int(userId))

	var products []models.Product
	var total float64

	for _, cartItem := range user.Cart {
		product, err := h.TransactionRepository.GetProductId(int(cartItem.ProductID))
		total += cartItem.Product.Price

		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		products = append(products, product)
	}

	// Cek apakah kode diskon valid
	discountPercentage := 0.0
	for _, discount := range discounts {
		if discount.Code == request.Discountcode {
			// Cek apakah diskon berlaku untuk kategori produk
			if discount.Category != "elektronik" {
				if !containsCategory(products, discount.Category) {
					return c.JSON(http.StatusBadRequest, dto.ErrorResult{
						Code:    http.StatusBadRequest,
						Message: fmt.Sprintf("Discount code %s is not applicable for the selected products", request.Discountcode),
					})
				}
			}

			// Cek apakah diskon berlaku untuk hari ini
			if len(discount.Days) > 0 && !containsDay(discount.Days, time.Now().Weekday()) {
				return c.JSON(http.StatusBadRequest, dto.ErrorResult{
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("Discount code %s is not applicable today", request.Discountcode),
				})
			}

			discountPercentage = discount.Discount
			break
		}
	}

	discountAmount := total * discountPercentage
	grandTotal := total - discountAmount

	transaction := models.Transaction{
		Status:             request.Status,
		Date:               time.Now(),
		UserID:             int(userId),
		Product:            products,
		Discountcode:       request.Discountcode,
		Discountpercentage: int(discountPercentage),
		Discountamount:     discountAmount,
		Total:              int(grandTotal),
		Qty:                len(user.Cart),
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	err = h.TransactionRepository.Delete(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func containsCategory(products []models.Product, category string) bool {
	for _, product := range products {
		if strings.ToLower(product.Category.Name) == strings.ToLower(category) {
			return true
		}
	}
	return false
}

func containsDay(days []time.Weekday, targetDay time.Weekday) bool {
	for _, day := range days {
		if day == targetDay {
			return true
		}
	}
	return false
}

func (h *handlerTransaction) UpdateTransaction(c echo.Context) error {
	request := new(transactionsdto.UpdateTransactionRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
func (h *handlerTransaction) Generatecsv(c echo.Context) error {
	transactions, err := h.TransactionRepository.Generatecsv()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// Membuka file CSV
	file, err := os.Create("data.csv")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Gagal membuat file CSV")
	}
	defer file.Close()

	// Membuat penulis CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Menulis data ke file CSV
	for _, transaction := range transactions {
		row := []string{strconv.Itoa(transaction.ID), transaction.User.Fullname, strconv.Itoa(transaction.Total), transaction.Date.Format("2006-01-02 15:04:05"), transaction.Status}
		err := writer.Write(row)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Gagal menulis data ke file CSV")
		}
	}

	// Mengatur header response
	c.Response().Header().Set(echo.HeaderContentType, "text/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=data.csv")

	// Mengirim file CSV sebagai response
	return c.File("data.csv")
}
