package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func JSON(c echo.Context, statusCode int, message string, data interface{}, success bool) error {
	returnedData := Response{
		Success:    success,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	return c.JSON(statusCode, returnedData)
}
