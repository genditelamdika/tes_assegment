package routes

import (
	"tour/handlers"
	"tour/pkg/middleware"
	"tour/pkg/mysql"
	"tour/repositories"

	"github.com/labstack/echo/v4"
)

func TripRoutes(e *echo.Group) {
	tripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(tripRepository)

	e.GET("/trips", h.FindTrips)
	e.GET("/trip/:id", h.GetTrip)
	e.POST("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip)))
	e.PATCH("/trip/:id", middleware.UploadFile(h.UpdateTrip))
	e.DELETE("/trip/:id", h.DeleteTrip)
}
