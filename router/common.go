package router

import (
	"email_auth/controller"
	"github.com/gin-gonic/gin"
)

func InitCommonRouter(router *gin.RouterGroup) {
	router.GET("/ping", controller.Ping)
}
