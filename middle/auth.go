package middle

import (
	"email_auth/model"
	"email_auth/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")

		if len(authHeader) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    model.ERROR,
				"message": "User identity is not authenticated",
			})

			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": model.ERROR,
				"msg":  "Jwt format error",
			})

			ctx.Abort()
			return
		}

		_, err := util.VerifyToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "Invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}

}
