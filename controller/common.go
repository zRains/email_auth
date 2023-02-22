package controller

import (
	"email_auth/model"
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	model.Result(model.Ok, true, nil, "操作成功", ctx)
}
