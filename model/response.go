package model

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Succeed bool   `json:"succeed"`
	Message string `json:"msg,omitempty"`
	Data    any    `json:"data,omitempty"`
}

const (
	Ok                  = 201
	NotOk               = 405
	Unauthorized        = 401
	Forbidden           = 403
	InternalServerError = 500
)

const (
	OkMsg                  = "操作成功"
	NotOkMsg               = "操作失败"
	UnauthorizedMsg        = "登录过期, 需要重新登录"
	LoginCheckErrorMsg     = "用户名或密码错误"
	ForbiddenMsg           = "无权访问该资源, 请联系网站管理员授权"
	InternalServerErrorMsg = "服务器内部错误"
)

const (
	ERROR   = 1
	SUCCESS = 0
)

func Result(code int, succeed bool, data any, msg string, ctx *gin.Context) {
	ctx.JSON(Ok, Response{
		code, succeed, msg, data,
	})
}

func ResultWithError(err error, ctx *gin.Context) {
	Result(ERROR, false, nil, err.Error(), ctx)
}
