package user

import (
	"github.com/gin-gonic/gin"
	"gserver/api"
	"gserver/pkg/e"
)

func (s Service) AddUser(c *gin.Context, req *api.User) (res *api.User, err e.Error) {
	res = &api.User{}
	u := convertToModel(req)
	err = s.dao.AddUser(u)
	res = u.ToAPI()
	return
}
