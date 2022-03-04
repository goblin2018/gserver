package v1

import (
	"gserver/api"
	"gserver/pkg/e"
	"gserver/pkg/r"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	req := &api.User{}
	er := c.ShouldBind(req)
	if er != nil {
		r.JSON(c, nil, e.InvalidParams.From(er))
		return
	}
}
