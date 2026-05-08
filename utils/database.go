package utils

import (
	"fmt"
	"path/filepath"
	"simple-block-api/config"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
)

func GetDB() *gorm.DB {
	// return getPostgresDB()
	return getSQLiteDB()
}

func getSQLiteDB() *gorm.DB {
	databaseGoDir := GetCurrentGoFileDir()
	dbFullPath := filepath.Join(databaseGoDir, "..", "sqlitedb", "test.db")
	db, err := gorm.Open(sqlite.Open(dbFullPath), &gorm.Config{
		Logger: logger.Default.LogMode(getLogMode()),
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}
func getLogMode() logger.LogLevel {

	var logMode logger.LogLevel

	switch strings.ToLower(config.Cfg.Server.LogLevel) {
	case "info":
		logMode = logger.Info // 输出所有SQL（开发用）
	case "warn":
		logMode = logger.Warn // 只输出警告
	case "error":
		logMode = logger.Error // 只输出错误
	default:
		logMode = logger.Silent // 关闭日志（生产推荐）
	}
	return logMode
}
func getPostgresDB() *gorm.DB {

	user := config.Cfg.Database.User
	password := config.Cfg.Database.Password
	dbname := config.Cfg.Database.Name

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		user,
		password,
		dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// 开发环境打开日志，生产可关闭
		Logger: logger.Default.LogMode(getLogMode()),
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 参数校验，防止恶意请求
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		if pageSize > 100 {
			pageSize = 100 // 设置单页最大条数
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
