package config

type appInfo struct {
	AppName string `mapstructure:"name" yaml:"name"`       // 应用名称
	Port    uint64 `mapstructure:"port" yaml:"port"`       // 端口号
	Prefix  string `mapstructure:"prefix" yaml:"prefix"`   // 接口前缀
	Version string `mapstructure:"version" yaml:"version"` // 版本号
	Monster bool   `mapstructure:"monster" yaml:"monster"` // 妖怪模式
}
