package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	controller "github.com/ranggabudipangestu/mvc-api/controllers"
	"github.com/ranggabudipangestu/mvc-api/lib/response"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return response.JSON(c, http.StatusOK, "Server is Running", nil)
	})

	//USERS
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/user/:id", controller.GetUserByIdController)
	e.DELETE("/user/:id", controller.DeleteUserController)
	e.PUT("/user/:id", controller.UpdateUserController)

	e.POST("/books", controller.CreateBookController)
	e.GET("/books", controller.GetBooksController)
	e.GET("/book/:id", controller.GetBookByIdController)
	e.PUT("/book/:id", controller.UpdateBookController)
	e.DELETE("/book/:id", controller.DeleteBookController)

	return e
}
