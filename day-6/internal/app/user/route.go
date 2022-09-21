package user

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/users", h.FindAll)
	g.POST("/user/:id", h.FindById)
}
