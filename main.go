package main

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/borntodie-new/todo-list-backup/config"
	"github.com/borntodie-new/todo-list-backup/utils"
)

var db *gorm.DB

func initDB() {
	var err error
	conf := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Fatalf("初始化Mysql连接失败：%s\n", err.Error())
	}
}

func main() {
	// 初始化日志对象
	utils.InitLogger()
	// 加载配置
	conf := config.GetConfig()
	zap.S().Info(conf.Redis.Host)
	zap.S().Info(conf.Mysql.Username)
}
