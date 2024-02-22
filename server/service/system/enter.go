package system

// service 层 入口
type ServiceGroup struct {
	InitDBService
	UserService
	JwtService
	CasbinService
	OperationRecordService
	AuthorityService
	MenuService
	BaseMenuService
	FileUploadAndDownloadService
}
