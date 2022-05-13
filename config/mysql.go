package config

import (
	"fmt"
)

// MySQL配置
type mysqlConfig struct {
	Host     string `mapstructure:"host" yaml:"host"`         // 主机
	Port     int    `mapstructure:"port" yaml:"port"`         // 端口
	User     string `mapstructure:"user" yaml:"user"`         // 用户名
	Password string `mapstructure:"password" yaml:"password"` // 密码
	Db       string `mapstructure:"db" yaml:"db"`             // 数据库名称
}

// GetDSN 返回 MySQL 连接字符串
func (c mysqlConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Db)
}
