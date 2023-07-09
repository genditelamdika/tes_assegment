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

func CartRoutes(e *echo.Group) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	// Middleware Rate Limiter
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

	e.GET("/carts", h.FindCart)
	e.GET("/cart/:id", h.GetCartById)
	//no
	e.GET("/cart/:id/user", h.GetUserByID)
	e.GET("/cart/:id/pending-products", h.GetPendingProducts)

	e.POST("/cart", middleware.Auth(h.CreateCart))
	e.DELETE("/cart/:id", middleware.Auth(h.DeleteCart))

	e.GET("/user/:id/cart", h.FindcartByUser)

}
