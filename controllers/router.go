package controllers

import (
	v1 "gserver/controllers/v1"
	"gserver/pkg/env"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	if env.ENV == env.Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	v1.LoadGroup(r)
	return r
}
