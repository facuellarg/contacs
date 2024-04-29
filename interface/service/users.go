package service

import "github.com/facuellarg/users/models"

type UserServiceI interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	DeleteUserById(id int) (bool, error)
	UpdateUser(user *models.User) (*models.User, error)
}
