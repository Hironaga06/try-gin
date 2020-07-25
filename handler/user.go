package handler

import (
	"net/http"
	"try_gin/request"
	"try_gin/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Users(db *gorm.DB) gin.HandlerFunc {
	users := service.FindAll(db)
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	}
}

func SetUser(db *gorm.DB) gin.HandlerFunc {
	req := &request.User{}
	return func(c *gin.Context) {
		c.Bind(req)
		service.Set(db, req)
		c.Status(http.StatusNoContent)
	}
}
