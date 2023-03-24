package middlers

import (
	"github.com/borntodie-new/todo-list-backup/config"
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	"github.com/borntodie-new/todo-list-backup/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader(constant.TokenOfHeaderKey)
		if tokenString == "" {
			ctx.JSON(http.StatusOK, resp.RespFailed(constant.TokenInvalidErr))
			ctx.Abort()
			return
		}
		customJWT := utils.NewCustomJWT([]byte(config.GetConfig().JWTConfig.SigningKey))
		claims, err := customJWT.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusOK, resp.RespFailed(err))
			ctx.Abort()
			return
		}
		ctx.Set(constant.IDOfContextKey, claims.ID)
		ctx.Set(constant.UsernameOfContextKey, claims.Username)
		ctx.Next()
	}
}
