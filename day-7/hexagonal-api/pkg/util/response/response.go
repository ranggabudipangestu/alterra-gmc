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

func JSON(c echo.Context, response *Response) error {
	return c.JSON(response.StatusCode, response)
}

func ReturnedData(success bool, statusCode int, message string, data interface{}) *Response {

	return &Response{
		Success:    success,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
