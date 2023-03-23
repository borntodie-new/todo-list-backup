package utils

import (
	"github.com/borntodie-new/todo-list-backup/config"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	ID       int64    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type CustomJWT struct {
	SigningKey []byte
}

func NewCustomJWT(signingKey []byte) *CustomJWT {
	return &CustomJWT{
		SigningKey: signingKey,
	}
}

// GenerateToken generate token by claims
func (j *CustomJWT) GenerateToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken parse token by token string
func (j *CustomJWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, constant.TokenInvalidErr
}

// RefreshToken refresh token by token string
func (j *CustomJWT) RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(config.GetConfig().RefreshTime) * time.Hour)) // refresh token expire time is 24 hour
		jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(j.SigningKey)
	}
	return "", constant.TokenInvalidErr
}
