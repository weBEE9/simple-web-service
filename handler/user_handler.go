package handler

import (
	"net/http"
	"strconv"
	"weBEE9/simple-web-service/service"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

// UserHandler ...
type UserHandler struct {
	Service service.UserService
}

func InitHandler(e *gin.Engine, s service.UserService) {
	h := UserHandler{
		Service: s,
	}

	e.GET("/v1/users", h.getAllUsers)
	e.GET("/v1/users/:id", h.getUserByID)
}

func (handler UserHandler) getAllUsers(c *gin.Context) {
	span := trace.SpanFromContext(c)
	defer span.End()
	span.SetName("getAllUsers")

	users, err := handler.Service.GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": http.StatusOK,
		"data": users,
	})
}

func (handler UserHandler) getUserByID(c *gin.Context) {
	span := trace.SpanFromContext(c)
	defer span.End()
	span.SetName("getUserByID")

	id := c.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	user, err := handler.Service.GetUserByID(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}
