package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ranggabudipangestu/mvc-api/config"
	"github.com/stretchr/testify/assert"
)

func initEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

func TestGetUsrControllers(t *testing.T) {
	testCases := []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{
		{
			name:                "success",
			path:                "/users",
			expectedStatus:      http.StatusOK,
			expectBodyStartWith: "{\"status\":\"success\",\"users\":[",
		},
	}

	e := initEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}
