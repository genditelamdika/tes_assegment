package routes

import (
	"tour/handlers"
	"tour/pkg/mysql"
	"tour/repositories"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Group) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(countryRepository)

	e.GET("/countrys", h.FindCountrys)
	e.GET("/country/:id", h.GetCountry)
	e.POST("/country", h.CreateCountry)
	e.PATCH("/country/:id", h.UpdateCountry)
	e.DELETE("/country/:id", h.DeleteCountry)
}
