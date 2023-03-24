package router

import (
	"github.com/borntodie-new/todo-list-backup/handler/user"
	"github.com/borntodie-new/todo-list-backup/middlers"
	"github.com/gin-gonic/gin"
)

var registerUserRouter = func(r *gin.RouterGroup, h user.HandlerService) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", h.Register)
		userRouter.POST("/login", h.Login)
		userRouter.GET("/user", middlers.Auth(), h.GetUser)
		userRouter.PUT("/user", middlers.Auth(), h.ChangePassword)
	}
}
