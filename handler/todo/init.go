package todo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler interface {
	AddTodo(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
	MGetTodo(ctx *gin.Context)
	DelTodo(ctx *gin.Context)
	UpdTodo(ctx *gin.Context)
}

type HandlerService struct {
	db *gorm.DB
}

var _ Handler = (*HandlerService)(nil)