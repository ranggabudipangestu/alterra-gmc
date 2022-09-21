package repository

import (
	"context"

	"github.com/ranggabudipangestu/hexagonal-api/internal/dto"
	"github.com/ranggabudipangestu/hexagonal-api/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]dto.ResponseUserDto, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.User, error)
	FindByEmail(ctx context.Context, email *string, username *string) (*model.User, error)
	Register(ctx context.Context, data *dto.RequestRegisterDto) error
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (r *user) FindAll(ctx context.Context, payload *dto.SearchGetRequest, p *dto.Pagination) ([]dto.ResponseUserDto, *dto.PaginationInfo, error) {

	var users []dto.ResponseUserDto
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.User{}).Find(&users)

	if payload.Search != "" {
		search := "%" + payload.Search + "%"
		query = query.Where("lower(name) LIKE ? or lower(email) Like ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(p)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, dto.CheckInfoPagination(p, count), err
}

func (r *user) FindByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	err := r.Db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *user) FindByEmail(ctx context.Context, email *string, username *string) (*model.User, error) {
	conn := r.Db.WithContext(ctx)

	var user model.User
	err := conn.Model(&model.User{}).Where("email = ? OR username= ?", email, username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *user) Register(ctx context.Context, data *dto.RequestRegisterDto) error {
	tx := r.Db.WithContext(ctx).Create(&model.User{
		Name:     data.Name,
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
	})
	err := tx.Error

	if err != nil {
		return err
	}
	return nil
}
