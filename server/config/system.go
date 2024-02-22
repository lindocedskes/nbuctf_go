package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Host          string `mapstructure:"host" json:"host" yaml:"host"`                               // 服务启动的主机地址
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                               // 服务端口
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`    // 路由前缀
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      // 数据库类型
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // 是否使用redis
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 是否使用多点登录拦截
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                   // 存储方式
}
