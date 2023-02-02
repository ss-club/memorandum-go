package serializer

import "gogogo/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	CreateAt int64  `json:"createAt"`
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}
}
