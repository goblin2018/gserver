package user

import (
	"gserver/api"
	"gserver/dao"
	"gserver/models"
)

type Service struct {
	dao *dao.Dao
}

func NewService() *Service {
	return &Service{dao.NewDao()}
}

func convertToModel(req *api.User) *models.User {
	return &models.User{
		Username: req.Username,
		Name:     req.Name,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Desc:     req.Desc,
		Email:    req.Email,
		Wechat:   req.Wechat,
	}
}
