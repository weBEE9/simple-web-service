package service

import (
	"weBEE9/simple-web-service/model"
	"weBEE9/simple-web-service/repository"
)

// DefaultUserService ...
type DefaultUserService struct {
	repo repository.UserRepository
}

// NewDefaultUserService ...
func NewDefaultUserService(repo repository.UserRepository) DefaultUserService {
	return DefaultUserService{
		repo: repo,
	}
}

// GetAllUsers ...
func (service DefaultUserService) GetAllUsers() ([]*model.User, error) {
	return service.repo.GetAllUsers()
}

// GetusbyID ...
func (service DefaultUserService) GetUserByID(id int64) (*model.User, error) {
	return service.repo.GetUserByID(id)
}

// GetusbyID ...
func (service DefaultUserService) GetUserByUsername(username string) (*model.User, error) {
	return service.repo.GetUserByUsername(username)
}

// GetusbyID ...
func (service DefaultUserService) CreateUser(u *model.User) error {
	return service.repo.CreateUser(u)
}

// GetusbyID ...
func (service DefaultUserService) UpdateUser(u *model.User) error {
	return service.repo.UpdateUser(u)
}

// GetusbyID ...
func (service DefaultUserService) DeleteUser(id int64) error {
	return service.repo.DeleteUser(id)
}
