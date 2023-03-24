package config

import (
	"go.uber.org/zap"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

type Redis struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database int    `mapstructure:"database"`
}

type JWTConfig struct {
	SigningKey  string `mapstructure:"signing_key"`
	ExpireTime  int    `mapstructure:"expire_time"`
	RefreshTime int    `mapstructure:"refresh_time"`
}

type EmailConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
}

type Config struct {
	*Server    `mapstructure:"server"`
	*Mysql     `mapstructure:"mysql"`
	*Redis     `mapstructure:"redis"`
	*JWTConfig `mapstructure:"jwt"`
	*EmailConfig `mapstructure:"email"`
}

var (
	configOnce sync.Once
	config     *Config
)

func GetConfig() *Config {
	configOnce.Do(func() {
		debug := os.Getenv("DEBUG")
		configFileName := "./config/conf-dev.yaml"
		if debug != "true" {
			configFileName = "./config/conf-pro.yaml"
		}
		v := viper.New()
		v.SetConfigFile(configFileName)
		if err := v.ReadInConfig(); err != nil {
			zap.S().Fatalf("读取配置文件失败：%s\n", err.Error())
		}
		config = &Config{}
		if err := v.Unmarshal(config); err != nil {
			zap.S().Fatalf("解析配置文件失败：%s\n", err.Error())
		}
	})
	return config
}
