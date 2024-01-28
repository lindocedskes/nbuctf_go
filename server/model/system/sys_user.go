package system

import (
	"github.com/gofrs/uuid/v5"
	"github.com/lindocedskes/model"
)

// 用户表
type SysUser struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	UUID            uuid.UUID         `json:"uuid" gorm:"index;comment:用户UUID"`                                            //在外部（如URL中）使用UUID来标识用户，以避免暴露用户的内部ID                                      //用户唯一id 主键                                          // 用户UUID
	Username        string            `json:"userName" gorm:"index;comment:用户登录名"`                                         // 用户登录名
	Password        string            `json:"-"  gorm:"comment:用户登录密码"`                                                    // 用户登录密码
	NickName        string            `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                   // 用户昵称
	SideMode        string            `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                 // 用户侧边主题
	HeaderImg       string            `json:"headerImg" gorm:"default:toadd;comment:用户头像"`                                 // 用户头像，TODO：头像默认上传
	BaseColor       string            `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                  // 基础颜色
	ActiveColor     string            `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                             // 活跃颜色
	AuthorityId     uint              `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                               // 用户角色ID
	Authority       SysAuthority      `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"` //外键，关联角色表
	Authorities     []SysAuthority    `json:"authorities" gorm:"many2many:sys_user_authority;"`                            //角色组
	Phone           string            `json:"phone"  gorm:"comment:用户手机号"`                                                 // 用户手机号
	Email           string            `json:"email"  gorm:"comment:用户邮箱"`                                                  // 用户邮箱
	Enable          int               `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                             //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
