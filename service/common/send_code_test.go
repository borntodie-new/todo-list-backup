package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var rd *redis.Client

func init() {
	rd = redis.NewClient(&redis.Options{
		Addr:     "192.168.226.130:6379",
		Password: "123456",
		DB:       0,
	})
	_, err := rd.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("失败")
		os.Exit(1)
	}
}

func TestSendCode(t *testing.T) {
	ctx := context.Background()
	email := "jasonborntodie@gmail.com"
	err := SendCode(email, ctx, rd)
	assert.Nil(t, err)
	result, err := rd.Get(ctx, "todo-list-backup-code-123456@qq.com").Result()
	assert.Nil(t, err)
	t.Log(result)
}
func TestSendCodeWithExpired(t *testing.T) {
	ctx := context.Background()
	email := "123456@qq.com"
	err := SendCode(email, ctx, rd)
	assert.Nil(t, err)
	time.Sleep(time.Second * 61)
	_, err = rd.Get(ctx, "todo-list-backup-code-123456@qq.com").Result()
	assert.Error(t, err)
}
