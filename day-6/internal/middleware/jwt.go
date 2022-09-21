package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ranggabudipangestu/hexagonal-api/internal/dto"
)

func JWTMiddleware(claims dto.JWTClaims, signingKey []byte) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &claims,
		SigningKey: signingKey,
	}
	return middleware.JWTWithConfig(config)
}
