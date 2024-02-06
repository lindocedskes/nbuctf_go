package request

import (
	"github.com/lindocedskes/model"
	"github.com/lindocedskes/model/system"
)

// 默认的菜单项,id为1,顶级菜单
func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		BaseModel: model.BaseModel{ID: 1},
		ParentId:  "0",
		Path:      "dashboard", //对应前端的路由path
		Name:      "dashboard",
		Component: "view/dashboard/index.vue", //对应的 Vue 组件的路径
		Sort:      1,                          //排序值
		Meta: system.Meta{ //菜单项的元数据
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
