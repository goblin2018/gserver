package dao

import (
	"gserver/api"
	"gserver/models"
	"gserver/pkg/e"
)

// AddUser 添加用户
func (d *Dao) AddUser(user *models.User) e.Error {
	var err error
	db := d.db
	if user.Phone == "" {
		db = db.Omit("phone")
	}
	err = db.Create(user).Error
	return e.DBError.From(err)
}

// UpdateUser 更新用户
func (d *Dao) UpdateUser(user *models.User) e.Error {
	err := d.db.Where("id = ?", user.ID).Updates(user).Error
	return e.DBError.From(err)
}

// DelUser 删除用户
func (d *Dao) DelUser(user *models.User) {
	d.db.Where("id = ?", user.ID).Delete(&models.User{})
}

// ListUser 获取用户列表
func (d *Dao) ListUser(opt *api.ListUserOpt) (users []*models.User, total int64, err e.Error) {
	er := d.db.Count(&total).Order("updated_at desc").Offset(opt.Offset).Limit(opt.Limit).Find(&users).Error
	err = e.DBError.From(er)
	return
}
