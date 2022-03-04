package dao

import (
	"gserver/models"

	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao() *Dao {
	return &Dao{models.GetDB()}
}
