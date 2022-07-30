package handler

import (
	"net/http"
	"strconv"
	"weBEE9/simple-web-service/service"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_users_count",
		Help: "The total number users calling",
	})

	successCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_users_success_count",
		Help: "The total number of users API success",
	})

	failCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_users_fail_count",
		Help: "The total number of users API failed",
	})
)

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

	opsProcessed.Inc()

	id := c.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		failCount.Inc()

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	user, err := handler.Service.GetUserByID(userID)
	if err != nil {

		failCount.Inc()

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	successCount.Inc()

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}
