package user

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserRequest struct {
	Id int64 `json:"id" binding:"required"`
}

func (h *Handler) GetUser(ctx *gin.Context) {
	req := new(GetUserRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	user, err := service.MGetUser([]int64{req.Id}, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccessWithData(user))
}
