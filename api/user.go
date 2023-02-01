package api

import (
	"github.com/gin-gonic/gin"
	"gogogo/service"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService

	if err := c.ShouldBind(&userRegisterService); err != nil {
		c.JSON(404, err)
		println(err)
	} else {
		res := userRegisterService.Register()
		c.JSON(200, res)
	}
}

func UserLogin(c *gin.Context) {
	var userLoginService service.UserService

	if err := c.ShouldBind(&userLoginService); err != nil {
		c.JSON(404, err)
		println(err)
	} else {
		res := userLoginService.Login()
		c.JSON(200, res)
	}
}
