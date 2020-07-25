package handler

import (
	"net/http"
	"try_gin/service"

	"github.com/gin-gonic/gin"
)

func Users() gin.HandlerFunc {
	users := service.All()
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	}
}
