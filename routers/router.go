package routers

import "github.com/gin-gonic/gin"

func InitRoute() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	return r
}
