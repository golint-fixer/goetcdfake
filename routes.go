package main

import (
	"github.com/labstack/echo"
)

func setRoutes(e *echo.Echo) {
	// Set key
	e.POST("/keys/:key/:value", addNewKeyHandler)

	// Get value of key
	e.GET("/keys/:key", getValueHandler)
}
