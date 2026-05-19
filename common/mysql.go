package common

import (
	"ai-workspace-backend/model"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// initialize sql by gorm, set connection pool

// global variable DB
var DB *gorm.DB

func InitMYSQL() error{
	dsn := "root:123456@tcp(127.0.0.1:3306)/ai?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect mysql: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get sql db: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&model.User{}); err != nil {
		return fmt.Errorf("auto migrate: %w", err)
	}
	DB = db
	return nil
}