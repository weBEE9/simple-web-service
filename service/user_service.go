package service

import "weBEE9/simple-web-service/model"

//go:generate mockgen -source=user_service.go -destination=../mock/user_service_mock.go -package=mock
// UserService ...
type UserService interface {
	GetAllUsers() ([]*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}
