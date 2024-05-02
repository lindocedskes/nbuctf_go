package system

//实现动态api
import (
	"github.com/lindocedskes/service/system"
)

const initOrderApi = system.InitOrderSystem + 1

func init() {
	//system.RegisterInit(initOrderApi, &initApi{})
}

//func (i initApi) InitializerName() string {
//	return sysModel.SysApi{}.TableName()
//}
//
//func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
//	db, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return ctx, system.ErrMissingDBContext
//	}
//	return ctx, db.AutoMigrate(&sysModel.SysApi{})
//}
//
//func (i *initApi) TableCreated(ctx context.Context) bool {
//	db, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return false
//	}
//	return db.Migrator().HasTable(&sysModel.SysApi{})
//}
//
//func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
//	db, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return ctx, system.ErrMissingDBContext
//	}
//	entities := getDefaultApi()
//	if err := db.Create(&entities).Error; err != nil {
//		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+"表数据初始化失败!")
//	}
//	next := context.WithValue(ctx, i.InitializerName(), entities)
//	return next, nil
//}
//
//func (i *initApi) DataInserted(ctx context.Context) bool {
//	db, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return false
//	}
//	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
//		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
//		return false
//	}
//	return true
//}
//
//func getDefaultApi() []sysModel.SysApi {
//	return []sysModel.SysApi{
//
//		// todo add more api
//		//{ApiGroup: "系统用户", Method: "DELETE", Path: "/user/deleteUser", Description: "删除用户"},
//		//{ApiGroup: "系统用户", Method: "POST", Path: "/user/admin_register", Description: "用户注册"},
//
//	}
//}
