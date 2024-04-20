package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base") // "/base" 开头的路由组
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)           // 登录
		baseRouter.POST("captcha", baseApi.Captcha)       // 验证码
		baseRouter.POST("register", baseApi.RegisterUser) // 用户注册账号
	}
	return baseRouter
}
