package todo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sync"
)

var (
	handler HandlerService
	handlerOnce sync.Once
)

type HandlerService interface {
	AddTodo(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
	MGetTodo(ctx *gin.Context)
	DelTodo(ctx *gin.Context)
	UpdTodo(ctx *gin.Context)
}

type Handler struct {
	db *gorm.DB
}

func NewHandlerService(db *gorm.DB) HandlerService {
	handlerOnce.Do(func() {
		handler = &Handler{
			db: db,
		}
	})
	return handler
}

var _ HandlerService = (*Handler)(nil)