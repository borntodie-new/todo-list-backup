package todo

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *HandlerService) GetTodo(ctx *gin.Context) {
	// userId from ctx get
	userId := ctx.MustGet(constant.IDOfContextKey).(int64)
	username := ctx.MustGet(constant.UsernameOfContextKey).(string)

	// todoId from url's params
	todoIdStr := ctx.Param("id")
	todoId, _ := strconv.ParseInt(todoIdStr, 10, 64)
	todo, err := service.GetTodo(username, userId, todoId, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccessWithData(todo))
}
