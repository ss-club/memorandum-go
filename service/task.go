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

func DeleteTask(id string) serializer.Response {
	var task model.Task
	err := model.DB.Model(model.Task{}).Where("id = ?", id).First(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除查询出错",
			Error:  err.Error(),
		}

	}
	err = model.DB.Delete(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

func (service *CreateTask) UpdateTask(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ?", id).First(&task)
	err := model.DB.Model(&task).Updates(model.Task{Title: service.Title, Content: service.Content, Status: service.Status}).Error
	//task.Content = service.Content
	//task.Status = service.Status
	//task.Title = service.Title

	//err := model.DB.Save(&task).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库更新出错",
			//Error:  id,
			Data: string(id),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "更新成功",
		Data:   task,
	}
}
