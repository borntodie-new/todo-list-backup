package model

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

var (
	db1 *gorm.DB
	td  *TodoDao
)

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
	db1, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln(err)
	}
	_ = db1.AutoMigrate(&Todo{})
	td = NewTodoDao(context.Background(), db1)
}

func TestTodoDao_CreateInstance(t *testing.T) {
	todo := &Todo{
		UserId:    2,
		Content:   "打豆豆",
		Completed: false,
	}
	err := td.CreateInstance(todo)
	assert.Nil(t, err)
}

func TestTodoDao_RetrieveInstances(t *testing.T) {
	var userId int64 = 1
	limit := 5
	offset := 0
	todos, err := td.RetrieveInstances(userId, offset, limit)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(todos))
}

func TestTodoDao_DeleteInstance(t *testing.T) {
	var userId int64 = 1
	var todoId int64 = 1

	err := td.DeleteInstance(userId, todoId)
	assert.Nil(t, err)
	todo, err := td.RetrieveInstance(userId, todoId)
	assert.Nil(t, err)
	assert.Equal(t, true, todo.Completed)
}

func TestTodoDao_UpdateInstance(t *testing.T) {
	var userId int64 = 1
	var todoId int64 = 1
	content := "完成Golang项目部分"
	err := td.UpdateInstance(userId, todoId, content)
	assert.Nil(t, err)

	todo, err := td.RetrieveInstance(userId, todoId)
	assert.Nil(t, err)
	assert.Equal(t, content, todo.Content)

}