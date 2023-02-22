package controller

import (
	"email_auth/model"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	model.OkWithMessage("pong", c)
}
