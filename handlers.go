package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func addNewKeyHandler(c echo.Context) error {
	key, value := c.Param("key"), c.Param("value")

	addKeyErr := addKey([]byte(key), []byte(value))
	if addKeyErr != nil {
		return c.String(http.StatusBadRequest, "false")
	}

	return c.String(http.StatusCreated, "true")
}

func getValueHandler(c echo.Context) error {
	key := c.Param("key")
	value := findKey([]byte(key))
	return c.String(http.StatusOK, string(value))
}
