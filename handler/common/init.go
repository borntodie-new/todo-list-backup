package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"sync"
)

type HandlerService interface {
	SendCode(ctx *gin.Context)
}

type Handler struct {
	rd *redis.Client
}

var (
	handler     HandlerService
	handlerOnce sync.Once
)

func NewHandlerService(rd *redis.Client) HandlerService {
	handlerOnce.Do(func() {
		handler = &Handler{
			rd: rd,
		}
	})
	return handler
}

var _ HandlerService = (*Handler)(nil)
