package http

import (
	"fmt"
	"myapp/config"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//InitHTTPServer Initialize HTTP Server
func InitHTTPServer() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType,
			echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost,
			http.MethodPut, http.MethodOptions},
		AllowCredentials: true,
	}))

	mountRoute(e)

	// Start server
	startErr := e.Start(fmt.Sprintf(":%d", config.Config.AppPort))
	if startErr != nil {
		fmt.Println(startErr)
		os.Exit(1)
	}
}
