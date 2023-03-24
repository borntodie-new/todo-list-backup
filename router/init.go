package router

import (
	"github.com/borntodie-new/todo-list-backup/handler/common"
	"github.com/borntodie-new/todo-list-backup/handler/user"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, rd *redis.Client) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	// register user module router
	userHandler := user.NewHandlerService(db, rd)
	registerUserRouter(api, userHandler)
	// register common modula router
	commonHandler := common.NewHandlerService(rd)
	registerCommonRouter(api, commonHandler)
	return r
}
