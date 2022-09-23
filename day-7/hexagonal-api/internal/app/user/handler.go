package user

import (
	"net/http"

	"github.com/ranggabudipangestu/hexagonal-api/internal/dto"
	"github.com/ranggabudipangestu/hexagonal-api/internal/factory"
	"github.com/ranggabudipangestu/hexagonal-api/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) FindAll(c echo.Context) error {
	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return response.JSON(c, response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil))
	}
	if err := c.Validate(payload); err != nil {
		return response.JSON(c, response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil))
	}

	res := h.service.Find(c.Request().Context(), payload)

	return response.JSON(c, res)
}

func (h *handler) FindById(c echo.Context) error {
	payload := new(dto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.JSON(c, response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil))
	}
	if err := c.Validate(payload); err != nil {
		return response.JSON(c, response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil))
	}

	res := h.service.FindByID(c.Request().Context(), payload)

	return response.JSON(c, res)
}
