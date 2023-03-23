package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJWT_GenerateToken(t *testing.T) {
	claims := &Claims{
		ID:       1,
		Username: "jason",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)), // 有效时间
			NotBefore: jwt.NewNumericDate(time.Now()),                       // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       // 签发时间
		},
	}
	customJWT := NewCustomJWT([]byte("secret"))
	token, err := customJWT.GenerateToken(claims)
	assert.Nil(t, err)
	t.Log(token)
}

func TestCustomJWT_ParseToken(t *testing.T) {
	claims := &Claims{
		ID:       1,
		Username: "jason",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)), // 有效时间
			NotBefore: jwt.NewNumericDate(time.Now()),                       // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       // 签发时间
		},
	}
	customJWT := NewCustomJWT([]byte("secret"))
	token, err := customJWT.GenerateToken(claims)
	assert.Nil(t, err)
	t.Log(token)
	result, err := customJWT.ParseToken(token)
	assert.Nil(t, err)
	assert.Equal(t, claims.ID, result.ID)
	assert.Equal(t, claims.Username, result.Username)
}
func TestCustomJWT_ParseTokenWithFalse(t *testing.T) {
	claims := &Claims{
		ID:       1,
		Username: "jason",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 10)), // 有效时间
			NotBefore: jwt.NewNumericDate(time.Now()),                       // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       // 签发时间
		},
	}
	customJWT := NewCustomJWT([]byte("secret"))
	token, err := customJWT.GenerateToken(claims)
	assert.Nil(t, err)
	t.Log(token)
	time.Sleep(time.Second * 11)
	_, err = customJWT.ParseToken(token)
	assert.Error(t, err)
}

func TestCustomJWT_RefreshToken(t *testing.T) {
	// 1. 签发token
	// 2. 即将过期
	// 3. 刷新token-》验证肯定成功
	claims := &Claims{
		ID:       1,
		Username: "jason",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 10)), // 有效时间
			NotBefore: jwt.NewNumericDate(time.Now()),                       // 生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       // 签发时间
		},
	}
	customJWT := NewCustomJWT([]byte("secret"))
	token, err := customJWT.GenerateToken(claims)
	assert.Nil(t, err)
	t.Log("第一次token：", token)
	t.Log("第一次token过期时间：", claims.RegisteredClaims.ExpiresAt)
	_, err = customJWT.ParseToken(token)
	assert.Nil(t, err)
	time.Sleep(time.Second * 5)
	refreshToken, err := customJWT.RefreshToken(token)
	assert.Nil(t, err)
	t.Log("第二次token：", refreshToken)
	result, err := customJWT.ParseToken(refreshToken)
	t.Log("第二次token过期时间：", result.RegisteredClaims.ExpiresAt)
	assert.Equal(t, claims.ID, result.ID)
	assert.Equal(t, claims.Username, result.Username)
}
