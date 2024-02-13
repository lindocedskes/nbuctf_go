package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

// @Summary   新建菜单记录
func (a *AuthorityMenuApi) CreateBaseMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = menuService.AddBaseMenu(menu)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}
