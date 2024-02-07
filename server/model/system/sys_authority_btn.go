package system

// 实现更细粒度的权限控制,表示每个角色对每个菜单的哪些按钮有操作权限。
type SysAuthorityBtn struct {
	AuthorityId      uint           `gorm:"comment:角色ID"`
	SysMenuID        uint           `gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:菜单按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:按钮详情"`
}
