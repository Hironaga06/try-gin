package main

import (
	"try_gin/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ngaut/log"
	"github.com/pkg/errors"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/", handler.Users())
	router.POST("/users", handler.SetUser())
	router.Run(":8080")
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
