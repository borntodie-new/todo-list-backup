package todo

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdTodoRequest struct {
	Content string `json:"content" binding:"required"`
}

func (h *HandlerService) UpdTodo(ctx *gin.Context) {
	req := new(UpdTodoRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	userId := ctx.MustGet(constant.IDOfContextKey).(int64)
	todoIdStr := ctx.Param("id")
	todoId, _ := strconv.ParseInt(todoIdStr, 10, 64)
	err := service.UpdateTodo(userId, todoId, req.Content, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess())
}
