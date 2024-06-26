package system

import (
	"context"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lindocedskes/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/*管理 CasbinRule 表的初始化器*/

type initApi struct{}

const initOrderCasbin = initOrderApi + 1

type initCasbin struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbin, &initCasbin{})
}

// 迁移表
func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i initCasbin) InitializerName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := getDefaultCasbinRule()
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitializerName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// 初始 CasbinRule 表数据 TODO add data
// Ptype: "p" 表示权限，Ptype: "g" 表示用户组，V0: "888" 表示角色ID，V1: "/user/getUserInfo" 表示接口路径，V2: "GET" 表示请求方法
func getDefaultCasbinRule() []adapter.CasbinRule {
	return []adapter.CasbinRule{
		//管理员的权限
		{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/init/checkdb", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/init/initdb", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/casbin/getCasbinRuleByAuthorityId", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/registerbyadmin", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setUserAuthorities", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/resetPassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/deleteAuthority", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/getAuthorityList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/updateAuthority", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/setDataAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/file/upload", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/file/getFileList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/file/deleteFile", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/file/editFileName", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/file/breakpointContinue", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/file/findFile", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/file/breakpointContinueFinish", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/file/removeChunk", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authorityBtn/getAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authorityBtn/setAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authorityBtn/canRemoveAuthorityBtn", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/game/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/game/submitflag", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/game/rank", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/game/submitScoreChart", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/game/filedownload", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/gameadmin/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/gameadmin/editquestion", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/gameadmin/createquestion", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/gameadmin/deletequestion", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/gameadmin/addfile", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/gameadmin/deletefiles", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/announcement/list", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/announcementadmin/create", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/announcementadmin/delete", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/announcementadmin/update", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/k8s/openContainer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/k8s/closeContainer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/k8s/checkContainer", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/k8s/closeAllContainer", V2: "GET"},

		//普通用户的权限
		{Ptype: "p", V0: "777", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "777", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "777", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/file/getFileList", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/file/findFile", V2: "GET"},
		{Ptype: "p", V0: "777", V1: "/game/list", V2: "GET"},
		{Ptype: "p", V0: "777", V1: "/game/submitflag", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/game/rank", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/game/submitScoreChart", V2: "GET"},
		{Ptype: "p", V0: "777", V1: "/game/filedownload", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/announcement/list", V2: "GET"},
		{Ptype: "p", V0: "777", V1: "/k8s/openContainer", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/k8s/closeContainer", V2: "POST"},
		{Ptype: "p", V0: "777", V1: "/k8s/checkContainer", V2: "POST"},
		//普通用户的权限
		{Ptype: "p", V0: "7771", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/user/changePassword", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/user/setSelfInfo", V2: "PUT"},
		{Ptype: "p", V0: "7771", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "7771", V1: "/menu/getMenu", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/file/getFileList", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/file/findFile", V2: "GET"},
		{Ptype: "p", V0: "7771", V1: "/game/list", V2: "GET"},
		{Ptype: "p", V0: "7771", V1: "/game/submitflag", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/game/rank", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/game/submitScoreChart", V2: "GET"},
		{Ptype: "p", V0: "7771", V1: "/game/filedownload", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/announcement/list", V2: "GET"},
		{Ptype: "p", V0: "7771", V1: "/k8s/openContainer", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/k8s/closeContainer", V2: "POST"},
		{Ptype: "p", V0: "7771", V1: "/k8s/checkContainer", V2: "POST"},
	}
}
