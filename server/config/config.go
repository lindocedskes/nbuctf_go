package config

// 配置文件中有的信息
type Server struct {
	//mapstructure，结构体标签（struct tag）
	//将这个结构体序列化为 JSON 或 YAML 格式时，应该使用 "system" 作为这个字段的键
	//反序列化时，配置文件对应键值对，赋值到该结构体，如System
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// oss 存储
	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	// 跨域配置
	//Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`

	//k3s
	K3s K3s `mapstructure:"k3s" json:"k3s" yaml:"k3s"`
}

// 结构体形式保存
