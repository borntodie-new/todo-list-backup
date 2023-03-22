package user

import (
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func (h *Handler) Register(ctx *gin.Context) {
	req := new(RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	if err := service.CreateUser(req.Username, req.Password, req.Email, req.Code, req.Avatar, ctx, h.db, h.rd); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess())
}
