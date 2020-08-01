package main

import (
	"os"
	"try_gin/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/juju/errors"
)

var (
	db *gorm.DB
)

func main() {
	defer db.Close()
	router := gin.Default()
	router.GET("/", handler.Users(db))
	router.POST("/users", handler.SetUser(db))
	router.Run(":" + os.Getenv("PORT"))
}

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(errors.Trace(err))
	}
}
