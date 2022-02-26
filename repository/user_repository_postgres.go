package repository

import (
	"fmt"

	"weBEE9/simple-web-service/model"

	"xorm.io/xorm"
)

// NewUserRepositoryPostgres ...
func NewUserRepositoryPostgres(e *xorm.Engine) UserRepository {
	return &UserRepositoryPostgres{
		engine: e,
	}
}

// UserRepositoryPostgres ...
type UserRepositoryPostgres struct {
	engine *xorm.Engine
}

// GetAllUsers ...
func (repo *UserRepositoryPostgres) GetAllUsers() ([]*model.User, error) {
	return getAllUsers(repo.engine)
}

// GetUserByID ...
func (repo *UserRepositoryPostgres) GetUserByID(id int64) (*model.User, error) {
	return getUserByID(repo.engine, id)
}

// GetUserByUsername ...
func (repo *UserRepositoryPostgres) GetUserByUsername(username string) (*model.User, error) {
	return getUserByUsername(repo.engine, username)
}

// CreateUser ...
func (repo *UserRepositoryPostgres) CreateUser(u *model.User) error {
	return createUser(repo.engine, u)
}

// UpdateUser ...
func (repo *UserRepositoryPostgres) UpdateUser(u *model.User, cols ...string) error {
	if len(cols) > 0 {
		return updateUserCols(repo.engine, u, cols...)
	}
	return updateAllUserCols(repo.engine, u)
}

func (repo *UserRepositoryPostgres) DeleteUser(id int64) error {
	return deleteUser(repo.engine, id)
}

func getAllUsers(e xorm.Interface) ([]*model.User, error) {
	users := make([]*model.User, 0)
	return users, e.Find(&users)
}

func getUserByID(e xorm.Interface, id int64) (*model.User, error) {
	user := &model.User{
		ID: id,
	}

	has, err := e.Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("user not found: %d", id)
	}

	return user, nil
}

func getUserByUsername(e xorm.Interface, usesrname string) (*model.User, error) {
	user := &model.User{
		Username: usesrname,
	}

	has, err := e.Get(user)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("user not found: %s", usesrname)
	}

	return user, nil
}

func createUser(e xorm.Interface, u *model.User) error {
	_, err := e.Insert(u)
	return err
}

func updateAllUserCols(e xorm.Interface, u *model.User) error {
	_, err := e.Update(u)
	return err
}

func updateUserCols(e xorm.Interface, u *model.User, cols ...string) error {
	_, err := e.Cols(cols...).Update(u)
	return err
}

func deleteUser(e xorm.Interface, id int64) error {
	_, err := e.ID(id).Delete(&model.User{})
	return err
}
