package core

import (
	"fmt"

	"github.com/Arthaslixin/FrozenThrone-go/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Logger.Warn("未配置mysql,取消gorm连接")
		return nil
	}

	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境显示所有的sql语句
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("[%s] mysql 连接失败", dsn))
		panic(err)
	}
	// sqlDB, _ := db.DB()
	// sqlDB.SetConnMaxLifetime(time.Hour * 4)                 // 连接最大服用时间,不能超过mysql的 wait_timeout
	return db
}
