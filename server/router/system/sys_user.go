package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("userbyadmin").Use(middleware.OperationRecord()) //管理员权限下
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("register", baseApi.Register)                     // 管理员注册用户账号
		userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限，通过修改主角色id todo ？？只能改当前用户的角色id
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 重置用户权限组，并第一个设为主角色
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户byid
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息byid
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("resetPassword", baseApi.ResetPassword)           //直接重置用户（byid）密码为默认值123456
		userRouter.POST("getUserList", baseApi.GetUserList)               // 分页获取用户列表
	}
	{
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo) // 获取自身信息 byUUid
	}
}
