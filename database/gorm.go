package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "sync"
	"cron2sql/structs"
)


var (
    db   *gorm.DB
    once sync.Once
)

func GetDB() *gorm.DB {
    if db == nil {
        panic("database is not initialized")
    }
    return db
}


// 初始化数据库
func InitDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // 自动迁移表结构
    err = db.AutoMigrate(&structs.Task{})
    if err != nil {
        return nil, err
    }

    return db, nil
}