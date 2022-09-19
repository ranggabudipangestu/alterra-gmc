package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ranggabudipangestu/mvc-api/lib/response"
	"github.com/ranggabudipangestu/mvc-api/lib/static"
	"github.com/ranggabudipangestu/mvc-api/models"
)

func GetBooksController(c echo.Context) error {
	books, err := static.GetBooks()

	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil, false)
	}

	return response.JSON(c, http.StatusOK, "Success", books, true)
}

func CreateBookController(c echo.Context) error {
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil, false)
	}

	err := static.CreateBook(book)

	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil, false)
	}

	return response.JSON(c, http.StatusOK, "success", "Book Successfully created", true)
}

func GetBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := static.GetBookById(id)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil, false)
	}

	return response.JSON(c, http.StatusOK, "success", data, true)
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := static.DeleteBook(id)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil, false)
	}

	return response.JSON(c, http.StatusOK, "success", "Data Successfully deleted", true)
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil, false)
	}

	err := static.UpdateBook(id, book)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil, true)
	}
	return response.JSON(c, http.StatusOK, "success", "User Successfully Updated", false)

}
