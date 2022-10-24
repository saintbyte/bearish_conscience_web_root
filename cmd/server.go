package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func setup(e *echo.Echo) {
	e.Debug = true
	e.Static("/static", "assets")
	e.File("/", "public/index.html")
	e.File("/favicon.ico", "assets/favicon.ico")
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func getPort() string {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "1323"
	}
	return httpPort
}

func main() {
	e := echo.New()
	setup(e)
	e.Logger.Fatal(e.Start(":" + getPort()))
}
