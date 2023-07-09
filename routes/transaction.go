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

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

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

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transaction/:id", h.GetTransaction)
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.PATCH("/transaction/:id", h.UpdateTransaction)
	e.DELETE("/transaction/:id", h.DeleteTransaction)

	e.GET("/generatecsv", h.Generatecsv)
}
