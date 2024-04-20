package system

// 路由入口
type RouterGroup struct {
	//todo toadd router
	BaseRouter
	InitRouter
	JwtRouter
	CasbinRouter
	UserRouter
	AuthorityRouter
	MenuRouter
	FileUploadAndDownloadRouter
	AuthorityBtnRouter
	GameRouter
}
