package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.GET("", getUser)
	return r
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello me",
	})
}
