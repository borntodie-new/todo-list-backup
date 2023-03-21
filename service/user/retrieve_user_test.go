package user

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

var db1 *gorm.DB

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
}

func TestRetrieveUser(t *testing.T) {
	var (
		username = "jason1"
		password = "jason123"
	)
	user, err := RetrieveUser(username, password, context.Background(), db1)
	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
	t.Log(user)
}
