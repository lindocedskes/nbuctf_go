package system

import (
	"github.com/gofrs/uuid/v5"
	"github.com/lindocedskes/model"
)

// 用户表
type SysUser struct {
	model.BaseModel `gorm:"embedded"` //嵌入字段，视为当前结构体的字段，否则不视为
	UUID            uuid.UUID         `json:"uuid" gorm:"index;comment:用户UUID"`                //在外部（如URL中）使用UUID来标识用户，以避免暴露用户的内部ID                                      //用户唯一id 主键                                          // 用户UUID
	Username        string            `json:"userName" gorm:"index;comment:用户登录名"`             // 用户登录名
	Password        string            `json:"-"  gorm:"comment:用户登录密码"`                        // 用户登录密码
	NickName        string            `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`       // 用户昵称
	SideMode        string            `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`     // 用户侧边主题
	HeaderImg       string            `json:"headerImg" gorm:"default:toadd;comment:用户头像"`     // 用户头像，TODO：头像默认上传
	BaseColor       string            `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`      // 基础颜色
	ActiveColor     string            `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"` // 活跃颜色
	AuthorityId     uint              `json:"authorityId" gorm:"default:888;comment:用户角色ID"`   // 用户角色ID，主要角色id
	//字段Authority，当你从数据库中查询SysUser时，GORM会自动填充Authority字段，使其包含与SysUser相关联的SysAuthority的信息
	//表示关联关系，设置外键，。foreignKey:AuthorityId 表示 SysUser 表中的 AuthorityId 字段是外键； references:AuthorityId 表示这个外键引用的是 SysAuthority 表中的 AuthorityId 字段； 2者自动匹配
	Authority SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"` //主要角色，根据 AuthorityId 相关联对应角色名，外键，关联角色表
	//多对多关系，外键关系是通过连接表维护，sys_user_authority是连接表，包含2个外键 用户 ID 和角色 ID，分别指向 SysUser 表和 SysAuthority 表
	//当你、向 SysUser.Authorities 切片添加一个 SysAuthority 对象时，sys_user_authority 表中添加一行，表示该用户拥有该角色
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"` //角色组，一个用户的所有角色
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"`                      // 用户手机号
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`                       // 用户邮箱
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`  //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
