package system

import "github.com/lindocedskes/model"

type SysUser struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	Username        string            `gorm:"primaryKey;unique;not null;comment:用户名;" json:"username"` // 用户名
	Password        string            `gorm:"not null;comment:用户密码;" json:"password"`                  // 用户密码
	Email           string            `gorm:"not null;comment:用户邮箱;" json:"email"`                     // 用户邮箱
	Nickname        string            `gorm:"not null;comment:用户昵称;" json:"nickname"`
	Roles           []SysRole         `gorm:"many2many:sys_user_roles;comment:用户角色;" json:"roles"`
	Enable          int               `gorm:"default:1;comment:账户状态;" json:"enable"`
	SecretKey       string            `gorm:"not null;comment:Jwt时效密钥;"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
