package config

// Nacos配置
type nacosConfig struct {
	Host             string `env:"NACOS_HOST"`                               // 主机
	Port             uint64 `env:"NACOS_PORT" envDefault:"8848"`             // 端口
	NamespaceId      string `env:"NACOS_NAMESPACE"`                          // 命名空间
	CenterConfigName string `env:"NACOS_CONFIG_NAME" envDefault:"gtest.yml"` // 外部配置文件名，多个以逗号隔开
}
