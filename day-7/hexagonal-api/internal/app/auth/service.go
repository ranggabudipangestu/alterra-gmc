package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ranggabudipangestu/hexagonal-api/internal/dto"
	"github.com/ranggabudipangestu/hexagonal-api/internal/factory"
	"github.com/ranggabudipangestu/hexagonal-api/internal/pkg/util"
	"github.com/ranggabudipangestu/hexagonal-api/internal/repository"
	"github.com/ranggabudipangestu/hexagonal-api/pkg/util/hash"
	"github.com/ranggabudipangestu/hexagonal-api/pkg/util/response"
)

type Service interface {
	Login(ctx context.Context, payload *dto.RequestLoginDto) *response.Response
	Register(ctx context.Context, payload *dto.RequestRegisterDto) *response.Response
}

type service struct {
	Repository repository.User
}

func NewService(f *factory.Factory) *service {
	return &service{f.UserRepository}
}

func (s *service) Login(ctx context.Context, payload *dto.RequestLoginDto) *response.Response {

	data, err := s.Repository.FindByEmail(ctx, &payload.Email, &payload.Email)
	if data == nil {
		return response.ReturnedData(false, http.StatusNotFound, "user not found", nil)
	}

	if checkHashedPassword := hash.CompareHashPassword(payload.Password, data.Password); !checkHashedPassword {
		return response.ReturnedData(false, http.StatusBadRequest, "email, username & password combination doesn't match", nil)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID, data.Username)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return response.ReturnedData(false, http.StatusInternalServerError, "failed to generate token", nil)

	}

	return response.ReturnedData(false, http.StatusOK, "success", &dto.ResponseLoginDto{Token: token})
}

func (s *service) Register(ctx context.Context, payload *dto.RequestRegisterDto) *response.Response {

	data, err := s.Repository.FindByEmail(ctx, &payload.Email, &payload.Username)
	if data != nil {
		return response.ReturnedData(false, http.StatusNotFound, fmt.Sprintf("Email %s OR username %s is already exists", payload.Email, payload.Username), nil)
	}

	hashedPassword, err := hash.HashPassword(payload.Password)
	if err != nil {
		return response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil)
	}
	payload.Password = hashedPassword
	err = s.Repository.Register(ctx, payload)
	if err != nil {
		return response.ReturnedData(false, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.ReturnedData(true, http.StatusOK, "Success", nil)

}
