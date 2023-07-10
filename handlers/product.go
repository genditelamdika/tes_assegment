package handlers

import (
	"fmt"
	productsdto "indocattes/dto/product"
	dto "indocattes/dto/result"
	"indocattes/models"
	"indocattes/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}
func (h *handlerProduct) FindProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	for i, p := range products {
		products[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: products})
}
func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	product.Image = path_file + product.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}
func (h *handlerProduct) CreateProduct(c echo.Context) error {
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	price, _ := strconv.Atoi(c.FormValue("price"))
	categoryid, _ := strconv.Atoi(c.FormValue("categoryid"))

	request := productsdto.CreateProductRequest{
		Name:        c.FormValue("name"),
		Price:       float64(price),
		Description: c.FormValue("description"),
		// Discountcode: c.FormValue("discountcode"),
		Image:      dataFile,
		CategoryID: categoryid,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	product := models.Product{
		Name:        request.Name,
		CategoryID:  request.CategoryID,
		Price:       request.Price,
		Description: request.Description,
		Image:       request.Image,
	}

	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	product, _ = h.ProductRepository.GetProduct(product.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	categoryid, _ := strconv.Atoi(c.FormValue("categoryid"))
	price, _ := strconv.Atoi(c.FormValue("price"))

	request := productsdto.UpdateProductRequest{
		Name:        c.FormValue("name"),
		CategoryID:  categoryid,
		Price:       float64(price),
		Description: c.FormValue("description"),
		Image:       dataFile,
	}

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Price != 0 {
		product.Price = request.Price
	}

	if request.CategoryID != 0 {
		product.CategoryID = request.CategoryID
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Image != "" {
		product.Image = request.Image
	}

	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsetrip(data)})
}
func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func convertResponsetrip(u models.Product) models.Product {
	return models.Product{
		ID:          u.ID,
		Name:        u.Name,
		CategoryID:  u.CategoryID,
		Price:       u.Price,
		Description: u.Description,
		Image:       u.Image,
	}
}
