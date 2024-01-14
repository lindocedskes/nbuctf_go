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
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//todo systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	//todo JWT
	//PrivateGroup := Router.Group(global.NBUCTF_CONFIG.System.RouterPrefix) // 需要鉴权的路由组
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//todo toadd router
	//systemRouter.InitApiRouter(PrivateGroup, PublicGroup)       // 注册功能api路由
	//systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
	//systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
	//systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
	//systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
	//systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
	//systemRouter.InitAutoCodeRouter(PrivateGroup)               // 创建自动化代码
	//systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
	//systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
	//systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)        // 自动化代码历史
	//systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
	//systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
	//systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)     // 字典详情管理
	//systemRouter.InitSysExportTemplateRouter(PrivateGroup)      // 导出模板
	//exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
	//exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	//}

	return Router
}
