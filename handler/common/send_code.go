package common

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SendCodeRequest struct {
	Email string `json:"email" binding:"required;email"`
}

func (h Handler) SendCode(ctx *gin.Context) {
	req := new(SendCodeRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	err := service.SendCode(req.Email, ctx, h.rd)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess())
}
