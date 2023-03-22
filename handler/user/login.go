package user

import (
	"fmt"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strings"

	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/user"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

func (h *Handler) Login(ctx *gin.Context) {
	req := new(LoginRequest)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	// handler code here
	key := fmt.Sprintf(constant.CodePrefix, req.Username)
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
	user, err := service.RetrieveUser(req.Username, req.Password, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	// TODO signed token here

	ctx.JSON(http.StatusOK, resp.RespSuccessWithData(user))
}
