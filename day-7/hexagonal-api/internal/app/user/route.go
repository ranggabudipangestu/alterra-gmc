package user

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("/users", h.FindAll)
	g.GET("/user/:id", h.FindById)
}
