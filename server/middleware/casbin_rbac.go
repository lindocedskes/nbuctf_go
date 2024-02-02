package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/service"
	"github.com/lindocedskes/utils"
	"strconv"
	"strings"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 进行基于角色的访问控制（RBAC）。请求与数据表中记录匹配 用户-角色-资源（请求路径）-请求方法，存在记录返回allow，否则deny。对表中记录操作根据casbin API
// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.NBUCTF_CONFIG.System.Env != "develop-casbin" { //develop-casbin模式不启用权限验证
			waitUse, _ := utils.GetClaims(c) //从请求中的token解析为claims
			//获取请求的PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.NBUCTF_CONFIG.System.RouterPrefix) //去掉前缀
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId)) //int转string，角色ID
			e := casbinService.Casbin()                   // 返回 生成的casbin执行器（根据模型+数据库连接）
			success, _ := e.Enforce(sub, obj, act)        //Casbin 的执行器，判断策略 中是否存在
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
