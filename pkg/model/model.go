package model

import (
	// GORM 的 Mysql 数据库驱动导入
	"goblog/pkg/logger"
	"goblog/pkg/types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:root@tcp(localhost:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})

	logger.LogError(err)

	return DB
}
