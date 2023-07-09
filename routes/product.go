package routes

import (
	"indocattes/handlers"
	"indocattes/pkg/middleware"
	"indocattes/pkg/mysql"
	"indocattes/repositories"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	limiter := tollbooth.NewLimiter(100, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			httpError := tollbooth.LimitByRequest(limiter, c.Response(), c.Request())
			if httpError != nil {
				return echo.NewHTTPError(httpError.StatusCode, httpError.Message)
			}
			return next(c)
		}
	})

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProduct)
	e.POST("/product", middleware.UploadFile(h.CreateProduct))
	e.PATCH("/product/:id", h.UpdateProduct)
	e.DELETE("/product/:id", h.DeleteProduct)
}
