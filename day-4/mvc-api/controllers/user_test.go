package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ranggabudipangestu/mvc-api/config"
	"github.com/ranggabudipangestu/mvc-api/lib/response"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

func TestGetUserController(t *testing.T) {
	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/users")
	if assert.NoError(t, GetUserController(c)) {
		var response response.Response
		json.Unmarshal([]byte(rec.Body.String()), &response)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, true, response.Success)
		assert.NotNil(t, response.Data)
	}
}

func TestGetUserByIdSuccess(t *testing.T) {
	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/user/7")

	if assert.NoError(t, GetUserByIdController(c)) {
		var res response.Response
		json.Unmarshal([]byte(rec.Body.String()), &res)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, res.Success)
		assert.NotNil(t, res.Data)
	}
}

func TestGetUserByIdFailed(t *testing.T) {
	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/user/1")
	if assert.NoError(t, GetUserByIdController(c)) {
		var res response.Response
		json.Unmarshal([]byte(rec.Body.String()), &res)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, false, res.Success)
		assert.Nil(t, res.Data)
	}
}
