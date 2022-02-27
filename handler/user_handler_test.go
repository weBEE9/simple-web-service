package handler

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"weBEE9/simple-web-service/mock"
	"weBEE9/simple-web-service/model"

	"github.com/appleboy/gofight/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

var (
	mockUserService *mock.MockUserService
	testHadler      UserHandler
)

// set up mock service
func setUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService = mock.NewMockUserService(ctrl)

	testHadler = UserHandler{
		Service: mockUserService,
	}
}

func TestGetAllUsers(t *testing.T) {
	setUp(t)

	router := gin.Default()
	router.GET("/users", testHadler.getAllUsers)
	fight := gofight.New()

	t.Run("should return 200 OK", func(t *testing.T) {
		testUsers := []*model.User{
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
		mockUserService.EXPECT().GetAllUsers().Return(testUsers, nil)

		fight.GET("/users").
			Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)

				data := r.Body.Bytes()
				users := gjson.GetBytes(data, "data")
				assert.Len(t, users.Array(), 2)
			})
	})

	t.Run("should return 500 Internal Server Error", func(t *testing.T) {
		mockUserService.EXPECT().GetAllUsers().Return(nil, errors.New("something went wrong"))

		fight.GET("/users").
			Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, r.Code)

				data := r.Body.Bytes()
				message := gjson.GetBytes(data, "message")
				assert.Equal(t, "something went wrong", message.String())
			})
	})
}

func TestGetUserByID(t *testing.T) {
	setUp(t)
	router := gin.Default()
	router.GET("/users/:id", testHadler.getUserByID)

	fight := gofight.New()

	t.Run("should return 200 OK", func(t *testing.T) {
		alice := &model.User{
			ID:       1,
			Username: "alice001",
			Name:     "Alice",
		}
		mockUserService.EXPECT().GetUserByID(int64(1)).Return(alice, nil)

		fight.GET("/users/1").
			Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)

				data := r.Body.Bytes()
				user := gjson.GetBytes(data, "data")
				v, ok := user.Value().(map[string]interface{})
				assert.True(t, ok)
				assert.Equal(t, alice.Name, v["name"])
				assert.Equal(t, alice.Username, v["username"])
			})
	})

	t.Run("should return error", func(t *testing.T) {
		mockUserService.EXPECT().GetUserByID(int64(2)).Return(nil, fmt.Errorf("user not found: %d", 2))

		fight.GET("/users/2").
			Run(router, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusInternalServerError, r.Code)

				data := r.Body.Bytes()
				message := gjson.GetBytes(data, "message")
				assert.Equal(t, "user not found: 2", message.String())
			})
	})
}
