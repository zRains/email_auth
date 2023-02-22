package router

import (
	"context"
	"email_auth/controller"
	"email_auth/service"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(ctx context.Context, userService service.UserService, router *gin.RouterGroup) {
	userController := controller.CreateUserController(ctx, userService)

	router.POST("/signup", userController.SignUp)
	router.POST("/signin", userController.SignIn)
}
