package todo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sync"
)

var (
	handler Handler
	handlerOnce sync.Once
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

func NewHandler(db *gorm.DB) Handler {
	handlerOnce.Do(func() {
		handler = &HandlerService{
			db: db,
		}
	})
	return handler
}

var _ Handler = (*HandlerService)(nil)