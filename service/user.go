package service

import (
	"gogogo/model"
	"gogogo/serializer"
	"gorm.io/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
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

	user.Username = service.UserName

	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{Status: 401, Msg: "生成密码错误"}
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return &serializer.Response{
			Status: 402,
			Msg:    "创建用户出错",
		}
	}

	return &serializer.Response{
		Status: 200,
		Msg:    "创建用户成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name", service.UserName).First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return serializer.Response{
				Status: 404,
				Msg:    "未找到用户",
			}
		}
		println(err)
		return serializer.Response{
			Status: 404,
			Msg:    "未找到用户其它错误",
		}
	}

	if !user.CheckPassword(service.Password) {
		return serializer.Response{Status: 500, Msg: "密码验证错误"}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "登录成功",
	}
}
