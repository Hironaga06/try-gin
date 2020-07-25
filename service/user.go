package service

import "try_gin/model"

func All() []*model.User {
	users := []*model.User{}

	users = append(users, &model.User{
		ID:   1,
		Name: "山田太郎",
	})

	return users
}
