package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
)

type JwtApi struct{}

// @Summary   jwt加入黑名单,并清除cookie中的token
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt) //加入黑名单
	if err != nil {
		global.GVA_LOG.Error("jwt作废失败", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		utils.ClearToken(c) //清除cookie中的token
		response.OkWithMessage("jwt作废成功", c)
	}
}
