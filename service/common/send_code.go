package common

import (
	"context"
	"fmt"
	"github.com/borntodie-new/todo-list-backup/constant"
	"github.com/borntodie-new/todo-list-backup/utils"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math/rand"
	"strings"
	"time"
)

type SendCodeFlow struct {
	// global data
	ctx context.Context
	rd  *redis.Client

	// request data
	Email string

	// temporary data
	code string

	// response data

}

func NewSendCodeFlow(email string, ctx context.Context, rd *redis.Client) *SendCodeFlow {
	return &SendCodeFlow{
		ctx:   ctx,
		rd:    rd,
		Email: email,
	}
}

func SendCode(email string, ctx context.Context, rd *redis.Client) error {
	return NewSendCodeFlow(email, ctx, rd).Do()
}

func (f *SendCodeFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.prepareData(); err != nil {
		return err
	}
	if err := f.packageData(); err != nil {
		return err
	}
	return nil
}

func (f *SendCodeFlow) checkParam() error {
	if f.Email == "" || !strings.Contains(f.Email, "@") {
		return constant.ParamErr
	}
	return nil
}

func (f *SendCodeFlow) prepareData() error {
	// send code to user's email
	value := genRandomCode(constant.CodeLength)
	zap.S().Infof("verify_code is %s\n", value)
	err := utils.NewEmailCodeService(f.Email, value, constant.EmailOfSubject).SendCode()
	if err != nil {
		return err
	}
	f.code = value
	return nil
}

func (f *SendCodeFlow) packageData() error {
	// save code to redis
	key := fmt.Sprintf(constant.CodePrefix, f.Email)
	_, err := f.rd.Set(f.ctx, key, f.code, constant.CodeExpires*time.Second).Result()
	return err
}

func genRandomCode(length int) string {
	rand.Seed(time.Now().Unix())
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
