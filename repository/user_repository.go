package repository

import "weBEE9/simple-web-service/model"

type UserRepository interface {
	GetUserByID(id int64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, cols ...string) error
	DeleteUser(id int64) error
}
