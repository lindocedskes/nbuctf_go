package system

type RouterGroup struct {
	//todo toadd router
	BaseRouter
	InitRouter
	JwtRouter
	CasbinRouter
	UserRouter
	AuthorityRouter
	MenuRouter
}
