package services

import "REST-API/models"

type UserService interface {
	CreatUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
