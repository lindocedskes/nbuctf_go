package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	systemRes "github.com/lindocedskes/model/system/response"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
)

type AuthorityApi struct{}

// @Summary   创建角色，具有默认casbin路由权限
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority, authBack system.SysAuthority
	var err error

	if err = c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if authBack, err = authorityService.CreateAuthority(authority); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
		return
	}
	err = casbinService.FreshCasbin()
	if err != nil {
		global.GVA_LOG.Error("创建成功，权限刷新失败。", zap.Error(err))
		response.FailWithMessage("创建成功，权限刷新失败。"+err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
}

// @Summary   删除角色
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	var err error
	if err = c.ShouldBindJSON(&authority); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 删除角色之前需要判断是否有用户正在使用此角色
	if err = authorityService.DeleteAuthority(&authority); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
		return
	}
	_ = casbinService.FreshCasbin() //刷新casbin策略
	response.OkWithMessage("删除成功", c)
}