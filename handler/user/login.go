package user

import (
	"github.com/borntodie-new/todo-list-backup/config"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"

	resp "github.com/borntodie-new/todo-list-backup/handler"
	service "github.com/borntodie-new/todo-list-backup/service/user"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(ctx *gin.Context) {
	req := new(LoginRequest)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(constant.ParamErr))
		return
	}
	// handler code here
	user, err := service.RetrieveUser(req.Username, req.Password, ctx, h.db)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	// signed token here
	conf := config.GetConfig()
	customJWT := utils.NewCustomJWT([]byte(config.GetConfig().JWTConfig.SigningKey))
	claims := &utils.Claims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().
				Add(time.Duration(conf.JWTConfig.ExpireTime) * time.Hour)), // 签名生效时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 签名生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签名生效时间
		},
	}
	token, err := customJWT.GenerateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.RespFailed(err))
		return
	}
	ctx.JSON(http.StatusOK, resp.RespSuccessWithData(gin.H{
		"user_id": user.ID,
		"token":   token,
	}))
}
