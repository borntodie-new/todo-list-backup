package todo

import (
	"github.com/borntodie-new/todo-list-backup/constant"
	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddTodoRequest struct {
	UserId  int64  `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (h *Handler) AddTodo(ctx *gin.Context) {
	req := new(AddTodoRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	err := service.CreateTodo(req.UserId, req.Content, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccess())
}
