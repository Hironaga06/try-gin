package service

import (
	"fmt"
	"try_gin/model"
	"try_gin/request"

	"github.com/jinzhu/gorm"
)

func FindAll(db *gorm.DB) *model.Users {
	users := &model.Users{}
	db.Find(users)

	return users
}

func Set(db *gorm.DB, req *request.User) {
	user := &model.User{
		ID:   req.ID,
		Name: req.Name,
	}
	fmt.Printf("set user: %+v\n", user)
}
