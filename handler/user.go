package handler

import (
	"net/http"
	"try_gin/request"
	"try_gin/service"

	"github.com/gin-gonic/gin"
)

func Users() gin.HandlerFunc {
	users := service.All()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	}
}

func SetUser() gin.HandlerFunc {
	req := &request.User{}
	return func(c *gin.Context) {
		c.Bind(req)
		service.Set(req)
		c.Status(http.StatusNoContent)
	}
}
