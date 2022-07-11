package http

import (
	"myapp/internal/migrate"
	"myapp/internal/movie"
	"net/http"

	"github.com/labstack/echo"
)

func mountRoute(e *echo.Echo) {
	v1 := e.Group("/v1")
	data := v1.Group("/data")

	e.GET("/", healthCheck)
	v1.GET("/movies", movie.Movie)
	data.GET("/titles", migrate.ReadJsonFileTitle)
	data.GET("/credits", migrate.ReadJsonFileCredit)
	v1.GET("/color", movie.Color)
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
