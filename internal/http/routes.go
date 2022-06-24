package http

import (
	"myapp/internal/migrate"
	"net/http"

	"github.com/labstack/echo"
)

func mountRoute(e *echo.Echo) {
	v1 := e.Group("/v1")
	data := v1.Group("/data")

	e.GET("/", healthCheck)
	data.GET("/titles", migrate.ReadJsonFileTitle)
	data.GET("/credits", migrate.ReadJsonFileCredit)
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
