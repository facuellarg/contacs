package repository

import "github.com/facuellarg/users/models"

type UserRepositoryI interface {
	CreateUser(user *models.User) (*models.User, error)
	GetuserById(id int) (*models.User, error)
	DeleteUserById(id int) (bool, error)
	UpdateUser(user *models.User) (*models.User, error)
}
