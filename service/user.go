package service

import (
	"gogogo/model"
	"gogogo/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"require,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"require,min=5,max=16"`
}

func (service *UserService) Register() *serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return &serializer.Response{
			Status: 400,
			Msg:    "用户已经存在",
		}
	}

	return &serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}
