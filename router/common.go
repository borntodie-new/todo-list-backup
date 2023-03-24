package router

import (
	"github.com/borntodie-new/todo-list-backup/handler/common"
	"github.com/gin-gonic/gin"
)

var registerCommonRouter = func(r *gin.RouterGroup, h common.HandlerService) {
	commonRouter := r.Group("/common")
	{
		commonRouter.POST("/send-code", h.SendCode) // send code
	}
}
