package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

var (
	handler     HandlerService
	handlerOnce sync.Once
)

func NewHandlerService(db *gorm.DB, rd *redis.Client) HandlerService {
	handlerOnce.Do(func() {
		handler = &Handler{db: db, rd: rd}
	})
	return handler
}

type HandlerService interface {
	Login(ctx *gin.Context)          // use RetrieveInstance to get user
	Register(ctx *gin.Context)       // use CreateInstance to create user
	GetUser(ctx *gin.Context)        // use RetrieveInstance to get user
	ChangePassword(ctx *gin.Context) // use UpdateInstanceOfPassword to update user
}
type Handler struct {
	db *gorm.DB
	rd *redis.Client
}

var _ HandlerService = (*Handler)(nil)
