package routes

import (
	"github.com/gin-gonic/gin"
	"gogogo/api"
	"gogogo/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
			authed.GET("tasks", api.ListTasks)
			authed.DELETE("task/:id", api.DeleteTask)
			authed.PUT("task", api.UpdateTask)
		}
	}

	return r
}
