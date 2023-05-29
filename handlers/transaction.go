package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	dto "tour/dto/result"
	transactiondto "tour/dto/transaction"
	"tour/models"
	"tour/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// var path_file = "http://localhost:5000/uploads/"

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
	for i, p := range transactions {
		transactions[i].Attachment = path_file + p.Attachment
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}
func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	transaction.Attachment = path_file + transaction.Attachment

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}
func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	tripid, _ := strconv.Atoi(c.FormValue("tripid"))
	counterqty, _ := strconv.Atoi(c.FormValue("counterqty"))
	total, _ := strconv.Atoi(c.FormValue("total"))

	request := transactiondto.CreateTransactionRequest{
		Counterqty: counterqty,
		Total:      total,
		Status:     c.FormValue("status"),
		Attachment: dataFile,
		TripID:     tripid,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transaction := models.Transaction{
		Counterqty: request.Counterqty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		TripID:     request.TripID,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	transaction, _ = h.TransactionRepository.GetTransaction(transaction.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
func (h *handlerTransaction) UpdateTransaction(c echo.Context) error {
	// request := new(tripdto.UpdateTripRequest)
	// if err := c.Bind(&request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	tripid, _ := strconv.Atoi(c.FormValue("tripid"))
	counterqty, _ := strconv.Atoi(c.FormValue("counterqty"))
	total, _ := strconv.Atoi(c.FormValue("total"))

	request := transactiondto.CreateTransactionRequest{
		Counterqty: counterqty,
		Total:      total,
		Status:     c.FormValue("status"),
		Attachment: dataFile,
		TripID:     tripid,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Counterqty != 0 {
		transaction.Counterqty = request.Counterqty
	}
	if request.Total != 0 {
		transaction.Total = request.Total
	}
	if request.Status != "" {
		transaction.Status = request.Status
	}
	if request.Attachment != "" {
		transaction.Attachment = request.Attachment
	}
	if request.TripID != 0 {
		transaction.TripID = request.TripID
	}
	// fmt.Println("", transaction)
	// if request.Trip != "" {
	// 	transaction.TripID = request.Trip
	// }

	data, err := h.TransactionRepository.UpdateTransaction(transaction, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func convertResponsetransaction(u models.Transaction) models.Transaction {
	return models.Transaction{
		ID:         u.ID,
		Counterqty: u.Counterqty,
		Total:      u.Total,
		Status:     u.Status,
		Attachment: u.Attachment,
		TripID:     u.TripID,
	}
}
