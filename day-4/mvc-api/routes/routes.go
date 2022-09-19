package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ranggabudipangestu/mvc-api/constants"
	controller "github.com/ranggabudipangestu/mvc-api/controllers"
	"github.com/ranggabudipangestu/mvc-api/lib/response"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return response.JSON(c, http.StatusOK, "Server is Running", nil, true)
	})

	//USERS

	e.POST("/users", controller.CreateUserController)
	e.GET("/books", controller.GetBooksController)
	e.GET("/book/:id", controller.GetBookByIdController)
	e.POST("/login", controller.LoginController)

	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.JWT_SECRET)))
	r.GET("/users", controller.GetUserController)
	r.GET("/user/:id", controller.GetUserByIdController)
	r.DELETE("/user/:id", controller.DeleteUserController)
	r.PUT("/user/:id", controller.UpdateUserController)
	r.POST("/books", controller.CreateBookController)
	r.PUT("/book/:id", controller.UpdateBookController)
	r.DELETE("/book/:id", controller.DeleteBookController)

	return e
}
