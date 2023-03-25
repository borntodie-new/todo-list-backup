package todo

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *HandlerService) MGetTodo(ctx *gin.Context) {
	// get userId from ctx
	userId := ctx.MustGet(constant.IDOfContextKey).(int64)
	todo, err := service.MGetTodo([]int64{userId}, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccessWithData(todo))
}