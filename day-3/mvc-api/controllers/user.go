package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ranggabudipangestu/mvc-api/lib/db"
	"github.com/ranggabudipangestu/mvc-api/lib/response"
	"github.com/ranggabudipangestu/mvc-api/middlewares"
	"github.com/ranggabudipangestu/mvc-api/models"
)

func GetUserController(c echo.Context) error {
	users, err := db.GetUsers()

	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}

	return response.JSON(c, http.StatusOK, "Success", users)
}

func CreateUserController(c echo.Context) error {
	var user *models.User

	if err := c.Bind(&user); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := db.CreateUsers(user)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}

	return response.JSON(c, http.StatusOK, "success", "User Successfully created")
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := db.GetUserById(id)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}

	return response.JSON(c, http.StatusOK, "success", data)
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := db.DeleteUser(id)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}

	return response.JSON(c, http.StatusOK, "success", "Data Successfully deleted")
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user *models.User

	if err := c.Bind(&user); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := db.UpdateUser(id, user)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	return response.JSON(c, http.StatusOK, "success", "User Successfully Updated")

}
func LoginController(c echo.Context) error {
	var user *models.LoginDto

	if err := c.Bind(&user); err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}

	userId, err := db.Login(user)

	if err != nil {
		return response.JSON(c, http.StatusBadRequest, "Failed to Login. Please check your email and password combination", nil)
	}

	token, err := middlewares.CreateToken(userId)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, "Failed to generate token", nil)
	}

	return response.JSON(c, http.StatusOK, "success", map[string]interface{}{"token": token})
}
