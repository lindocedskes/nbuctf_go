package system

import (
	"errors"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/request"
	"github.com/lindocedskes/model/system"
	"gorm.io/gorm"
	"strconv"
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

// @description: 按角色id获取动态菜单树存储在menus , treeMap[] ，system.SysMenu类型，键：父菜单的ID，值：包含所有子菜单的切片。绑定了角色id和 Btns
func (menuService *MenuService) GetMenuTree(authorityId uint) (menus []system.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"] //获取顶级菜单
	//遍历获取子菜单，存储在menus
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree) //遍历各个顶级菜单
	}
	return menus, err
}

// @description: 获取菜单树 treeMap[] ，system.SysMenu类型，键：父菜单的ID，值：包含所有子菜单的切片。绑定了角色id和 Btns
func (menuService *MenuService) getMenuTreeMap(authorityId uint) (treeMap map[string][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	var btns []system.SysAuthorityBtn
	treeMap = make(map[string][]system.SysMenu)

	var SysAuthorityMenus []system.SysAuthorityMenu
	//查询authorityId对应的所有SysAuthorityMenu记录
	err = global.NBUCTF_DB.Where("sys_authority_authority_id = ?", authorityId).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string
	//得到所有菜单id
	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuId)
	}
	//查询出MenuIds中所有ID对应的SysBaseMenu
	err = global.NBUCTF_DB.Where("id in (?)", MenuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}
	//SysBaseMenu转换为SysMenu（绑定更多信息）添加到allMenu切片中
	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityId: authorityId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}
	//sys_base_menu_btns 表中按角色id查询，记录包含菜单id，菜单按钮id，按钮详情
	err = global.NBUCTF_DB.Where("authority_id = ?", authorityId).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint) //2级键
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]uint)
		}
		//遍历按钮，对按钮的每个菜单创建2级map，值为角色id
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityId
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.ID]                                //将每个菜单的按钮设置为与菜单ID对应的按钮
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v) //将每个菜单添加到其父菜单的子菜单列表
	}
	return treeMap, err
}

// @description: 获取子菜单
func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId] //传入父菜单id
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// @description: 获取基础路由菜单树
func (menuService *MenuService) GetBaseMenuTree() (menus []system.SysBaseMenu, err error) {
	treeMap, err := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

// @description: 获取路由总树map，查询SysBaseMenu表
func (menuService *MenuService) getBaseMenuTreeMap() (treeMap map[string][]system.SysBaseMenu, err error) {
	var allMenus []system.SysBaseMenu
	treeMap = make(map[string][]system.SysBaseMenu)
	err = global.NBUCTF_DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// @description: 获取基础子菜单
func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[string][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// @description: 分页获取基础menu列表
func (menuService *MenuService) GetInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var menuList []system.SysBaseMenu
	treeMap, err := menuService.getBaseMenuTreeMapByPage(limit, offset)
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap) //子菜单
	}
	err = global.NBUCTF_DB.Model(&system.SysBaseMenu{}).Where("parent_id = ?", "0").Count(&total).Error
	return menuList, total, err
}

// @description: 获取路由总树map，查询SysBaseMenu表 -分页
func (menuService *MenuService) getBaseMenuTreeMapByPage(limit int, offset int) (treeMap map[string][]system.SysBaseMenu, err error) {
	var allMenus []system.SysBaseMenu
	treeMap = make(map[string][]system.SysBaseMenu)
	err = global.NBUCTF_DB.Limit(limit).Offset(offset).Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}
