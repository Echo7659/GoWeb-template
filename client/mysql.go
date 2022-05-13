package client

import (
	"GoWeb-template/config"

	"github.com/Echo7659/elog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQL *gorm.DB

func InitMySQLClient() {
	// 创建连接对象
	mysqlConfig := mysql.Config{
		DSN:                     config.Scd.MySQL.GetDSN(),
		DontSupportRenameIndex:  true, // 重命名索引时采用删除并新建的方式
		DontSupportRenameColumn: true, // 用 `change` 重命名列
	}
	conn, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{Logger: elog.DefaultGormLogger()})
	if err != nil {
		elog.Say.Panic("初始化MySQL连接失败, 错误信息: %v", err)
	} else {
		elog.Say.Debug("MySQL连接成功")
	}
	MySQL = conn

	//db, err := conn.DB()
	//db.SetMaxIdleConns(10)           // 用于设置连接池中空闲连接的最大数量
	//db.SetMaxOpenConns(100)          // 用于设置数据库连接的最大打开数量
	//db.SetConnMaxLifetime(time.Hour) // 设置连接的最大存活时间
}
