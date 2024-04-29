package repository

import (
	"sync"

	"github.com/facuellarg/users/models"
	"github.com/google/uuid"
)

type UserRepositoryMap struct {
	mutex sync.Mutex
	Users map[int]*models.User
}

var (
	ErrorUserNotFound = models.Error{Code: 404, Message: "user not found"}
	ErrorUserExists   = models.Error{Code: 409, Message: "user already exists"}
	ErrInternalServer = models.Error{Code: 500, Message: "internal server error"}
)

var once sync.Once

var userRepositoryMap UserRepositoryMap

func GetInstance() *UserRepositoryMap {
	once.Do(func() {
		userRepositoryMap = *newUserRepositoryMap()
	})
	return &userRepositoryMap
}

func newUserRepositoryMap() *UserRepositoryMap {
	return &UserRepositoryMap{
		mutex: sync.Mutex{},
		Users: make(map[int]*models.User),
	}
}

func (ur *UserRepositoryMap) CreateUser(user *models.User) (*models.User, error) {
	if _, ok := ur.Users[user.Id]; ok {
		return nil, ErrorUserExists
	}
	//generate unique id using uuid
	newId := uuid.New()
	user.Id = int(newId.ID())
	ur.registerUser(user)

	return user, nil
}

func (ur *UserRepositoryMap) GetuserById(id int) (*models.User, error) {
	return ur.Users[id], nil
}

func (ur *UserRepositoryMap) DeleteUserById(id int) (bool, error) {
	_, ok := ur.Users[id]
	ur.mutex.Lock()
	delete(ur.Users, id)
	ur.mutex.Unlock()
	return ok, nil
}

func (ur *UserRepositoryMap) UpdateUser(user *models.User) (*models.User, error) {
	if _, ok := ur.Users[user.Id]; !ok {
		return nil, ErrorUserNotFound
	}
	ur.registerUser(user)
	return user, nil
}

func (ur *UserRepositoryMap) registerUser(user *models.User) {
	ur.mutex.Lock()
	ur.Users[user.Id] = user
	ur.mutex.Unlock()
}
