package todo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

var db4 *gorm.DB

func init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	dsn := "root:123456@tcp(192.168.226.130:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db4, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func TestUpdateTodo(t *testing.T) {
	var (
		ctx           = context.Background()
		userId  int64 = 5
		todoId  int64 = 6
		content       = "用户5修改第一条todo任务"
	)
	err := UpdateTodo(userId, todoId, content, ctx, db4)
	assert.Nil(t, err)
}