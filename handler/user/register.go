package user

import (
	"fmt"
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func (h *Handler) Register(ctx *gin.Context) {
	req := new(RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	// handler code here
	key := fmt.Sprintf(constant.CodePrefix, req.Email)
	cacheCode, err := h.rd.Get(ctx, key).Result()
	if err == redis.Nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.CodeExpiresErr))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.CodeIncorrectErr))
		return
	}
	if strings.ToLower(cacheCode) != strings.ToLower(req.Code) {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.CodeIncorrectErr))
		return
	}
	if err = service.CreateUser(req.Username, req.Password, req.Email, req.Avatar, ctx, h.db); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess())
}
