package db

import (
	"github.com/ranggabudipangestu/mvc-api/config"
	"github.com/ranggabudipangestu/mvc-api/models"
)

func GetUsers() ([]models.UserDto, error) {

	var users []models.UserDto
	if err := config.DB.Model(&models.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUsers(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id int) (*models.UserDto, error) {
	var user models.UserDto
	if err := config.DB.Model(&models.User{}).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(id int) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(id int, data *models.User) error {

	user := &models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
