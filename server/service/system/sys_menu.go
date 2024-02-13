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

// @description: 添加基础路由
func (menuService *MenuService) AddBaseMenu(menu system.SysBaseMenu) error {
	if !errors.Is(global.NBUCTF_DB.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.NBUCTF_DB.Create(&menu).Error
}

// @description: 为角色增加menu树，前端会根据菜单树生成导航菜单
func (menuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, authorityId uint) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth) //连接表sys_authority_menus 添加一条关联记录
	return err
}
