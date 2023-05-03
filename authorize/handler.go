package authorize

import (
	"github.com/labstack/echo"
)

func ValidateHandlePrivateScoped(ctx echo.Context) bool {
	claims := parseClaims(ctx)
	return claims.HasScope("read:messages")
}
