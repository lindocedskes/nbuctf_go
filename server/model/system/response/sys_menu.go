package response

import "github.com/lindocedskes/model/system"

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus"`
}
type SysBaseMenusResponse struct {
	Menus []system.SysBaseMenu `json:"menus"`
}
