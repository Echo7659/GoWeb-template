package config

var Scd systemConfigData
var Nacos nacosConfig

// 配置信息
type systemConfigData struct {
	Admin appInfo     `mapstructure:"admin" yaml:"admin"` // 系统配置-Admin
	Api   appInfo     `mapstructure:"api" yaml:"api"`     // 系统配置-Api
	MySQL mysqlConfig `mapstructure:"mysql" yaml:"mysql"` // MySQL配置
	Redis redisConfig `mapstructure:"redis" yaml:"redis"` // Redis配置
}
