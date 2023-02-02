package service

import (
	"gogogo/model"
	"gogogo/serializer"
	"gorm.io/gorm"
	"time"
)

type CreateTask struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=1,max=100"`
	Content string `form:"content" json:"content" binding:"max=200"`
	Status  int    `form:"status" json:"status"`
}

func (service *CreateTask) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建备忘录额失败",
			Error:  err.Error()}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建备忘录成功",
		Data:   serializer.BuildTask(task),
	}
}

func ListTasks(id uint) serializer.Response {
	var tasks []model.Task
	err := model.DB.Model(model.Task{}).Where("uid = ?", id).Find(&tasks).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Response{
				Status: 404,
				Msg:    "未找到相关备忘录",
			}
		} else {
			return serializer.Response{
				Status: 500,
				Msg:    "查询备忘录出错",
			}
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "查询成功",
		Data:   serializer.BuildTasks(tasks),
	}

}
