package controller

import (
	"context"
	"email_auth/model"
	"email_auth/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type UserControllerImpl struct {
	ctx         context.Context
	userService service.UserService
}

func CreateUserController(ctx context.Context, userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		ctx,
		userService,
	}
}

func (this UserControllerImpl) SignUp(ctx *gin.Context) {
	var signUpUser model.SignUpRequired

	if err := ctx.ShouldBindJSON(&signUpUser); err != nil {
		model.FailWithMessage(err.Error(), ctx)

		return
	}

	// Store user into database
	dbResult, err := this.userService.SignUp(&model.User{
		Name:     signUpUser.Name,
		Email:    signUpUser.Email,
		Password: signUpUser.Password,
	})

	if err != nil {
		model.FailWithMessage(err.Error(), ctx)

		return
	}

	model.OkWithData(dbResult, ctx)
}

func (this UserControllerImpl) SignIn(ctx *gin.Context) {
	var signInUser model.SignInRequired

	if err := ctx.ShouldBindJSON(&signInUser); err != nil {
		model.FailWithMessage(err.Error(), ctx)

		return
	}

	dbUser, err := this.userService.SignIn(&signInUser)

	if err != nil {
		model.FailWithMessage(err.Error(), ctx)

		return
	}

	model.OkWithData(dbUser, ctx)
}
