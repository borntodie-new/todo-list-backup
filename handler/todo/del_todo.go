package todo

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) DelTodo(ctx *gin.Context) {
	// get userId from ctx
	userId := ctx.MustGet(constant.IDOfContextKey).(int64)
	// get todoId from url's params
	todoIdStr := ctx.Param("id")
	todoId, _ := strconv.ParseInt(todoIdStr, 10, 64)
	err := service.DeleteTodo(userId, todoId, ctx, h.db)
	if err != nil{
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess())
}