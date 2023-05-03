package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HandleHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func HandlePrivate(c echo.Context) error {
	return c.JSON(http.StatusOK, "Private")
}

func HandlePrivateScoped(c echo.Context) error {
	return c.JSON(http.StatusOK, "Private Scoped")
}
