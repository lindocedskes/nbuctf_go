package system

// 系统api层的代码，注册 对 service 层的调用 相关的结构体和变量进行封装
import "github.com/lindocedskes/service"

type ApiGroup struct {
	BaseApi
	DBApi
	JwtApi
}

var ( // 为了在 service/system 包中方便地使用这些服务，简写服务路径
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
