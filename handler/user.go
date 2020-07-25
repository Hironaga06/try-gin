package handler

import (
	"net/http"
	"try_gin/service"

	"github.com/gin-gonic/gin"
)

func Hoge() gin.HandlerFunc {
	hoges := service.All()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, hoges)
	}
}
