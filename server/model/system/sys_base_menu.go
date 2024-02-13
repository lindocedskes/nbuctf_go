package system

import (
	"github.com/lindocedskes/model"
)

// 存储系统的菜单结构以实现动态菜单、菜单的权限管理
// 构建导航菜单：Name 字段用作菜单项的文本（唯一），Path 字段前端路由的路径，Component 字段指定了点击菜单项后加载的vue页面组件，Sort 菜单列表中的排序。Meta 导航菜单中显示额外的信息
type SysBaseMenu struct {
	model.BaseModel
	MenuLevel     uint                                       `json:"-"`
	ParentId      string                                     `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path          string                                     `json:"path" gorm:"comment:路由path"`        // 路由path
	Name          string                                     `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden        bool                                       `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component     string                                     `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort          int                                        `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	Meta          `json:"meta" gorm:"embedded;comment:附加属性"` // 附加属性
	SysAuthoritys []SysAuthority                             `json:"authoritys" gorm:"many2many:sys_authority_menus;"` // 角色-菜单 连接表
	Children      []SysBaseMenu                              `json:"children" gorm:"-"`                                // 子菜单
	Parameters    []SysBaseMenuParameter                     `json:"parameters"`                                       // 路由参数
	MenuBtn       []SysBaseMenuBtn                           `json:"menuBtn"`                                          // 基础菜单按钮
}

type Meta struct {
	ActiveName  string `json:"activeName" gorm:"comment:高亮菜单"`
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
}

type SysBaseMenuParameter struct {
	model.BaseModel
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`            // 地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`            // 地址栏携带参数的值
}

func (SysBaseMenu) TableName() string {
	return "sys_base_menus"
}
