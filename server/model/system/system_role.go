package system

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	RoleName string `gorm:"unique;not null;comment:角色名称;" json:"role_name"`
}

func (SysRole) TableName() string {
	return "sys_roles"
}
