package system

type SysAuthorityMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	AuthorityId string `json:"-" gorm:"comment:角色ID;column:sys_authority_authority_id"`
}
