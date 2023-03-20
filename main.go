package main

import (
	"github.com/borntodie-new/todo-list-backup/config"
	"github.com/borntodie-new/todo-list-backup/utils"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志对象
	utils.InitLogger()
	// 加载配置
	conf := config.GetConfig()
	zap.S().Info(conf.Redis.Host)
	zap.S().Info(conf.Mysql.Username)
}
