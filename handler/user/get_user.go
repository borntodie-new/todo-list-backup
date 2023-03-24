package user

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUser(ctx *gin.Context) {
	userId := ctx.MustGet(constant.IDOfContextKey).(int64)
	user, err := service.MGetUser([]int64{userId}, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccessWithData(user))
}
