package system

import (
	"errors"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService) //指针，共享实例

// 根据用户的权限 ID 来查询相关的菜单 ID，然后再查询菜单，若无则默认路由设为 404
func (menuService *MenuService) UserAuthorityDefaultRouter(user *system.SysUser) {
	var menuIds []string // 菜单ID
	// 在用户角色菜单表中，根据用户的权限 ID 来查询相关的菜单 ID
	err := global.NBUCTF_DB.Model(&system.SysAuthorityMenu{}).Where("sys_authority_authority_id = ?", user.AuthorityId).Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am system.SysBaseMenu
	err = global.NBUCTF_DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
