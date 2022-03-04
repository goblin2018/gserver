package models

import (
	"gserver/api"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);comment:登录用户名;uniqueIndex"`
	Name     string `gorm:"type:varchar(255);comment:姓名"`
	Nickname string `gorm:"type:varchar(255);comment:昵称"`
	Phone    string `gorm:"type:varchar(255);comment:手机号码;uniqueIndex"`
	Avatar   string `gorm:"type:varchar(255);comment:头像"`
	Desc     string `gorm:"type:varchar(255);comment:自我介绍"`
	Email    string `gorm:"type:varchar(255);comment:邮箱"`
	Wechat   string `gorm:"type:varchar(255);comment:微信"`
}

func (u *User) ToAPI() *api.User {
	a := &api.User{
		ID:       u.ID,
		Username: u.Username,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Name:     u.Name,
		Phone:    u.Phone,
		Email:    u.Email,
		Wechat:   u.Wechat,
		Desc:     u.Desc,
	}
	return a
}
