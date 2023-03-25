package router

import (
	"github.com/borntodie-new/todo-list-backup/handler/todo"
	"github.com/borntodie-new/todo-list-backup/middlers"
	"github.com/gin-gonic/gin"
)

var registerTodoRouter = func(r *gin.RouterGroup, h todo.HandlerService) {
	todoRouter := r.Group("/todo").Use(middlers.Auth())
	{
		todoRouter.POST("/add-todo", h.AddTodo)
		todoRouter.GET("/get-todo/:id", h.GetTodo)
		todoRouter.GET("/mget-todo", h.MGetTodo)
		todoRouter.DELETE("/del-todo/:id", h.DelTodo)
		todoRouter.PUT("/upd-todo/:id", h.UpdTodo)
	}
}
