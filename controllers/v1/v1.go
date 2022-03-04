package v1

import "github.com/gin-gonic/gin"

func LoadGroup(r *gin.Engine) {
	v1 := r.Group("/v1")

	user := v1.Group("/user")
	{
		user.POST("")
	}

}
