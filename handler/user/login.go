package user

import (
	"github.com/gin-gonic/gin"
	"net/http"

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
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
	}
	user, err := service.RetrieveUser(req.Username, req.Password, req.Code, ctx, h.db, h.rd)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
	}
	ctx.JSON(http.StatusOK, resp.RespSuccessWithData(user))
}
