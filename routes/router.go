package routes

import (
	"github.com/gin-gonic/gin"
	"gogogo/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
	}

	return r
}
