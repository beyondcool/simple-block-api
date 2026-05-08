package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var Cfg *Config

func InitViper() {
	// 这里可以添加 Viper 的初始化代码，例如设置配置文件路径、读取配置等
	// 设置配置文件名称（不含扩展名）
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 添加配置文件搜索路径
	viper.AddConfigPath(".")
	viper.AddConfigPath(GetCurrentGoFileDir())

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
		panic(err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("解析配置失败: %v", err)
		panic(err)
	}

	log.Println("✅ 配置文件加载成功")
}

func GetCurrentGoFileDir() string {
	// 获取当前 .go 文件所在的目录
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("获取当前文件路径失败")
	}
	return filepath.Dir(filename)
}
