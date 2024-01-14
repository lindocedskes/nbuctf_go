package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                            // 服务启动的主机地址
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                            // 服务端口
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 路由前缀
}
