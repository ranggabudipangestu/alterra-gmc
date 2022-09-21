package http

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/ranggabudipangestu/hexagonal-api/internal/app/auth"
	"github.com/ranggabudipangestu/hexagonal-api/internal/factory"
	"github.com/ranggabudipangestu/hexagonal-api/pkg/util"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	auth.NewHandler(f).Route(v1.Group("/auth"))
}
