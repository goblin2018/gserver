package v1

import (
	"gserver/api"
	"gserver/pkg/e"
	"gserver/pkg/r"
	"gserver/services/user"

	"github.com/gin-gonic/gin"
)

var userService = user.NewService()

func AddUser(c *gin.Context) {
	req := &api.User{}
	er := c.ShouldBind(req)
	if er != nil {
		r.JSON(c, nil, e.InvalidParams.From(er))
		return
	}
	res, err := userService.AddUser(c, req)
	r.JSON(c, res, err)
}
