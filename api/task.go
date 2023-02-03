package api

import (
	"github.com/gin-gonic/gin"
	"gogogo/service"
	"gogogo/utils"
)

func CreateTask(c *gin.Context) {
	createService := service.CreateTask{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "创建失败")
	}
}

func ListTasks(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := service.ListTasks(claim.Id)
	c.JSON(200, res)
}

func DeleteTask(c *gin.Context) {
	res := service.DeleteTask(c.Param("id"))
	c.JSON(200, res)
}

func UpdateTask(c *gin.Context) {
	updateService := service.CreateTask{}
	if err := c.ShouldBind(&updateService); err == nil {
		res := updateService.UpdateTask(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(404, "参数错误")
	}
}
