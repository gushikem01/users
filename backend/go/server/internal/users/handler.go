package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gushikem01/users-go/pkg/uusers"
	"go.uber.org/zap"
)

type UserHandler struct {
	log *zap.Logger
	us  *uusers.Service
}

// NewUsers
func NewUsers(log *zap.Logger, us *uusers.Service) *UserHandler {
	return &UserHandler{log, us}
}

func (uh *UserHandler) AddRoute(r *gin.RouterGroup) {
	u := r.Group("/users")
	{
		u.GET("/", uh.get)
	}
}

func (uh *UserHandler) get(c *gin.Context) {
	u, err := uh.us.Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "400 bad request")
	}
	// fmt.Println(u)
	c.JSON(http.StatusOK, u)
}
