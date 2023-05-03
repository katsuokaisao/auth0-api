package handler

import (
	"net/http"

	"github.com/katsuokaisao/auth0-api/authorize"
	"github.com/labstack/echo"
)

func HandleHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}

func HandlePrivate(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Private")
}

func HandlePrivateScoped(ctx echo.Context) error {
	ok := authorize.ValidateHandlePrivateScoped(ctx)
	if !ok {
		return ctx.JSON(http.StatusForbidden, "Forbidden")
	}

	return ctx.JSON(http.StatusOK, "Private Scoped Pass")
}
