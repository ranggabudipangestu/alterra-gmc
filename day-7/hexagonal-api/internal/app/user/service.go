package user

import (
	"context"
	"net/http"

	"github.com/ranggabudipangestu/hexagonal-api/internal/dto"
	"github.com/ranggabudipangestu/hexagonal-api/internal/factory"
	"github.com/ranggabudipangestu/hexagonal-api/internal/repository"
	"github.com/ranggabudipangestu/hexagonal-api/pkg/constant"
	"github.com/ranggabudipangestu/hexagonal-api/pkg/util/response"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) *response.Response
	FindByID(ctx context.Context, payload *dto.ByIDRequest) *response.Response
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) *response.Response {

	users, info, err := s.UserRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		if err == constant.RecordNotFound {
			return response.ReturnedData(false, http.StatusNotFound, err.Error(), nil)
		}
		return response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil)
	}

	result := new(dto.SearchGetResponse[dto.ResponseUserDto])
	result.Datas = users
	result.PaginationInfo = *info

	return response.ReturnedData(true, http.StatusOK, "Success", result)
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) *response.Response {
	var result *dto.ResponseUserDto

	data, err := s.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {

		if err == constant.RecordNotFound {
			return response.ReturnedData(false, http.StatusNotFound, err.Error(), nil)
		}
		return response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil)
	}

	result = &dto.ResponseUserDto{
		ID:       int(data.ID),
		Name:     data.Name,
		Email:    data.Email,
		Username: data.Username,
	}

	return response.ReturnedData(true, http.StatusOK, "success", result)
}
