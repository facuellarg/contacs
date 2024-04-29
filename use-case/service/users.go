package service

import (
	"github.com/facuellarg/users/models"
	"github.com/facuellarg/users/use-case/repository"
)

type UserService struct {
	UserRepository repository.UserRepositoryI
}

func NewUserService(userRepository repository.UserRepositoryI) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	return us.UserRepository.CreateUser(user)
}

func (us *UserService) GetUserById(id int) (*models.User, error) {
	return us.UserRepository.GetuserById(id)
}

func (us *UserService) DeleteUserById(id int) (bool, error) {
	return us.UserRepository.DeleteUserById(id)
}

func (us *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return us.UserRepository.UpdateUser(user)
}
