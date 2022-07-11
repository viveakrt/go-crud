package movie

import (
	"myapp/pkg/growthbook"
	"net/http"

	"github.com/labstack/echo"
)

func Movie(c echo.Context) error {
	id := c.Request().Header.Get("id")
	gb := growthbook.Gbook(id)
	if gb.Feature("movie").On {
		title := fetchMovie()
		return c.JSON(http.StatusOK, title)
	}
	return c.JSON(http.StatusNotFound, "404")

}

func Color(c echo.Context) error {
	var color Colour
	id := c.Request().Header.Get("id")
	color.ID = id
	gb := growthbook.Gbook(id)
	if gb.Feature("color").On {
		color.Color = "Blue"
		return c.String(http.StatusOK, id+",blue")
	}
	color.Color = "Purple"
	return c.String(http.StatusOK, id+",purple")
}
