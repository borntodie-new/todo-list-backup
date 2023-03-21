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
	db *gorm.DB
	ud *UserDao
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln(err)
	}
	_ = db.AutoMigrate(&User{})
	ud = NewUserDao(context.Background(), db)
}

func TestUserDao_CreateInstance(t *testing.T) {
	user := &User{
		Username:  "tank",
		Password:  "tank123",
		Email:     "tank@gmail.com",
		Avatar:    "https://www.baidu.com/avatar/default.png",
	}
	err := ud.CreateInstance(user)
	assert.Nil(t, err)
}

func TestUserDao_RetrieveInstances(t *testing.T) {
	uIds := []int64{1, 2}
	users, err := ud.RetrieveInstances(uIds)
	assert.Nil(t, err)
	assert.Equal(t, len(uIds), len(users))
}

func TestUserDao_RetrieveInstance(t *testing.T) {
	username := "jason"
	user, err := ud.RetrieveInstance(username)
	assert.Nil(t, err)
	assert.Equal(t, username, user.Username)
}

func TestUserDao_UpdateInstanceOfPassword(t *testing.T) {
	username := "jason"
	password := "jasonNB"
	err := ud.UpdateInstanceOfPassword(username, password)
	assert.Nil(t, err)
	user, err := ud.RetrieveInstance(username)
	assert.Nil(t, err)
	assert.Equal(t, password, user.Password)
}