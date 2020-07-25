package main

import (
	"net/http"

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
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8080")
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=postgres password= sslmode=disable")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
