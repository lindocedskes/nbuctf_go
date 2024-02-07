package system

import "time"

// 角色表
type SysAuthority struct {
	CreatedAt       time.Time       // 创建时间
	UpdatedAt       time.Time       // 更新时间
	DeletedAt       *time.Time      `sql:"index"`
	AuthorityId     uint            `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string          `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        *uint           `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`             // 角色ID-数据权限ID 连接表，角色具有哪些数据权限的（通过角色ID拥有其他角色来表示）
	Children        []SysAuthority  `json:"children" gorm:"-"`                                                   // 子角色
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`                         // 角色-菜单 通过 sys_authority_menus连接表建立其AuthorityId与sys_base_menu表主键id的关系
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`                              // 用户-角色 连接表
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`                 // 默认菜单(默认dashboard)
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
