package system

// 系统api层的代码，注册 对 service 层的调用 相关的结构体和变量进行封装
import "github.com/lindocedskes/service"

// control 层 api入口
type ApiGroup struct { //api 方法调用入口,control 层
	BaseApi
	DBApi
	JwtApi
	CasbinApi
	AuthorityApi
	AuthorityMenuApi
	FileUploadAndDownloadApi
}

var ( // 为了在 service/system 包中方便地使用这些服务，简写服务路径
	userService                  = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService                = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	jwtService                   = service.ServiceGroupApp.SystemServiceGroup.JwtService
	casbinService                = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	authorityService             = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	menuService                  = service.ServiceGroupApp.SystemServiceGroup.MenuService
	baseMenuService              = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	fileUploadAndDownloadService = service.ServiceGroupApp.SystemServiceGroup.FileUploadAndDownloadService
)
