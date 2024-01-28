package config

type Mysql struct { //嵌套结构体, ,inline ,squash将嵌套结构体的字段视为顶层结构体的字段，
	GeneralDB `yaml:",inline" mapstructure:",squash"` // 内联（inline），相当于把 GeneralDB 的字段都“压平”到这里
}

// 连接数据库所需要的所有信息，如用户名、密码、主机名、端口号、数据库名等
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
