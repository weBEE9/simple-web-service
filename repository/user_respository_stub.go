package repository

import (
	"fmt"
	"strings"

	"weBEE9/simple-web-service/model"
)

// NewUserRepositoryStub ...
func NewUserRepositoryStub() *UserRepositoryStub {
	Users := []*model.User{
		{
			ID:       1,
			Username: "alice001",
			Name:     "Alice",
		},
		{
			ID:       2,
			Username: "bob001",
			Name:     "Bob",
		},
	}
	return &UserRepositoryStub{
		users: Users,
	}
}

// UserRepositoryStub
type UserRepositoryStub struct {
	users []*model.User
}

// GetAllUsers ...
func (repo *UserRepositoryStub) GetAllUsers() ([]*model.User, error) {
	return repo.users, nil
}

// GetUserByID ...
func (repo *UserRepositoryStub) GetUserByID(id int64) (*model.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found: %d", id)
}

// GetUserByUsername ...
func (repo *UserRepositoryStub) GetUserByUsername(username string) (*model.User, error) {
	for _, user := range repo.users {
		if strings.EqualFold(user.Username, username) {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found: %v", username)
}

// CreateUser ...
func (repo *UserRepositoryStub) CreateUser(u *model.User) error {
	// fake auto incr ID
	autoIncrID := int64(len(repo.users) + 1)
	u.ID = autoIncrID

	repo.users = append(repo.users, u)
	return nil
}

// UpdateUser ...
func (repo *UserRepositoryStub) UpdateUser(u *model.User, cols ...string) error {
	for i, user := range repo.users {
		if user.ID == u.ID {
			repo.users[i].Name = u.Name
			repo.users[i].Username = u.Username
			return nil
		}
	}

	return fmt.Errorf("user not found: %d", u.ID)
}

func (repo *UserRepositoryStub) DeleteUser(id int64) error {
	for i, user := range repo.users {
		if user.ID == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user not found: %d", id)
}
