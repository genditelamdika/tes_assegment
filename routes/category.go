package routes

import (
	"indocattes/handlers"
	"indocattes/pkg/mysql"
	"indocattes/repositories"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Group) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	e.GET("/categorys", h.FindCategorys)
	e.GET("/category/:id", h.GetCategory)
	e.POST("/category", h.CreateCategory)
	e.PATCH("/category/:id", h.UpdateCategory)
	e.DELETE("/category/:id", h.DeleteCategory)
}
