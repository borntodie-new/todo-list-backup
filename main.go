package main

import (
	"context"
	"fmt"
	"github.com/borntodie-new/todo-list-backup/router"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"

	"github.com/borntodie-new/todo-list-backup/config"
	"github.com/borntodie-new/todo-list-backup/utils"
)

var (
	db *gorm.DB
	rd *redis.Client
)

func initDB() {
	var err error
	conf := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Infof("init mysql connect fail：%s\n", err.Error())
		os.Exit(1)
	}
}

func initRD() {
	conf := config.GetConfig()
	rd = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: config.GetConfig().Redis.Password,
		DB:       0,
	})
	_, err := rd.Ping(context.Background()).Result()
	if err != nil {
		zap.S().Infof("init redis connect fail：%s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	// 初始化日志对象
	utils.InitLogger()
	// 加载配置
	conf := config.GetConfig()
	initDB()
	initRD()
	engine := router.InitRouter(db, rd)
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	zap.S().Info(s.ListenAndServe())
}
