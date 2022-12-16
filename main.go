package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// entry point for SPA react app served by echo
	e.Static("/", "public")
	e.GET("/", func(c echo.Context) error {
		return c.File("public/index.html")

	})
	e.Logger.Fatal(e.Start(":8080"))
}
