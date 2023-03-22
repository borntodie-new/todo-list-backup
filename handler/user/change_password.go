package user

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChangeRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
	RepPassword string `json:"rep_password" binding:"required;eqfield=NewPassword"`
}

func (h *Handler) ChangePassword(ctx *gin.Context) {
	req := new(ChangeRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	err := service.UpdateUser(req.OldPassword, req.NewPassword, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess())
}
