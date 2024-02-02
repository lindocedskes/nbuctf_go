package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/docs"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/middleware"
	"github.com/lindocedskes/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers() *gin.Engine {
	if global.NBUCTF_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //设置为发布模式
	}

	Router := gin.New()        //创建一个gin实例
	Router.Use(gin.Recovery()) //恢复中间件，从任何panic恢复，如果有panic的话，会写入500响应码
	if global.NBUCTF_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger()) //日志中间件，记录每个 HTTP 请求的详细信息
	}

	systemRouter := router.RouterGroupApp.System //系统路由
	//exampleRouter := router.RouterGroupApp.Example //测试路由

	//静态文件设置
	Router.Static("/favicon.ico", "./dist/favicon.ico")                                              // 设置静态文件目录的路由。两个参数：路由路径和文件系统路径
	Router.Static("/assets", "./dist/assets")                                                        // dist里面的静态资源
	Router.StaticFile("/", "./dist/index.html")                                                      // 前端网页入口页面，设置单个静态文件的路由。两个参数：路由路径和文件系统路径，带文件后缀
	Router.StaticFS(global.NBUCTF_CONFIG.Local.Path, http.Dir(global.NBUCTF_CONFIG.Local.StorePath)) //设置一个静态文件系统的（路由路径，文件系统路径）

	//swagger ui 自动生成接口文档
	docs.SwaggerInfo.BasePath = global.NBUCTF_CONFIG.System.RouterPrefix                                               // 设置路由前缀
	Router.GET(global.NBUCTF_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // 注册swagger路由
	global.GVA_LOG.Info("register swagger handler")

	// 跨域
	Router.Use(middleware.Cors())              // 直接放行全部跨域请求
	global.GVA_LOG.Info("use middleware cors") //日志记录

	PublicGroup := Router.Group(global.NBUCTF_CONFIG.System.RouterPrefix) // 不需要鉴权的路由组
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitDbRouter(PublicGroup)   // 初始化数据库数据
		systemRouter.InitBaseRouter(PublicGroup) // 基础登录功能路由 不做鉴权 todo 普通用户9528注册功能
	}
	PrivateGroup := Router.Group(global.NBUCTF_CONFIG.System.RouterPrefix) // 只在PrivateGroup路由组 下使用JWTAuth中间件和CasbinHandler中间件（每次【请求 之前 c.next() 之后】 执行中间件函数）
	//中间件链，中间件按照它们被添加的顺序依次执行，每个中间件都可以决定是否继续执行后续的中间件和处理函数
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()) //每次请求执行，先jwt有效性验证并获取claims，再casbin权限验证
	{
		//	//todo toadd router
		//systemRouter.InitApiRouter(PrivateGroup, PublicGroup)       // 存储API信息，更容易地实现权限管理
		systemRouter.InitJwtRouter(PrivateGroup)    // jwt相关路由
		systemRouter.InitCasbinRouter(PrivateGroup) // casbin_rule 的更新和查询
		systemRouter.InitUserRouter(PrivateGroup)   // 管理员注册用户路由
		//	//systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
		//systemRouter.InitMenuRouter(PrivateGroup) // 注册menu路由，数据库存储菜单信息可以：以根据用户的权限、角色或其他条件来定制菜单；为每个菜单项分配一个或多个角色，然后只允许具有这些角色的用户看到或访问这个菜单项

		//	//systemRouter.InitSystemRouter(PrivateGroup)                 // system配置相关路由：查改重载系统
		//	//systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
		//	//exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		//	//exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		//
	}

	return Router
}
