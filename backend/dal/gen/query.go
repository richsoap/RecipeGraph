package gen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化SQLite数据库连接并设置为默认连接
func InitDB(dbPath string) (*gorm.DB, error) {
	// 确保数据库目录存在
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// 配置GORM日志
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 200,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 设置为默认连接
	SetDefault(db)

	return db, nil
}

// InitDBWithDefault 使用默认路径初始化数据库
func InitDBWithDefault() (*gorm.DB, error) {
	defaultPath := "./database/recipe.db"
	return InitDB(defaultPath)
}
