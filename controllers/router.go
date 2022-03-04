package controllers

import (
	"fmt"
	v1 "gserver/controllers/v1"
	"gserver/pkg/env"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	if env.ENV == env.Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	fmt.Println(env.ENV)

	v1.LoadGroup(r)
	return r
}
