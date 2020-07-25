package service

import (
	"fmt"
	"try_gin/model"
	"try_gin/request"
)

func All() []*model.User {
	users := []*model.User{}

	users = append(users, &model.User{
		ID:   1,
		Name: "山田太郎",
	})

	return users
}

func Set(req *request.User) {
	user := &model.User{
		ID:   req.ID,
		Name: req.Name,
	}
	fmt.Printf("set user: %+v\n", user)
}
