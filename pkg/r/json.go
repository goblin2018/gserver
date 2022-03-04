package r

import (
	"gserver/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Abort(c *gin.Context, err e.Error) {
	c.Abort()
	JSON(c, nil, err)
}

func JSON(c *gin.Context, data interface{}, err e.Error) {
	c.JSON(http.StatusOK, gin.H{
		"code": err.Code(),
		"msg":  err.Error(),
		"data": data,
	})
}
