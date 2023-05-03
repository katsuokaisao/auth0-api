package authorize

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo"
)

func parseToken(ctx echo.Context) *CustomClaims {
	token := ctx.Request().Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	claims := token.CustomClaims.(*CustomClaims)

	return claims
}

func ValidateHandlePrivateScoped(ctx echo.Context) bool {
	claims := parseToken(ctx)
	return claims.HasScope("read:messages")
}
