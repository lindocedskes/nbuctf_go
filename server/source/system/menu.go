package system

import (
	"context"
	. "github.com/lindocedskes/model/system"
	"github.com/lindocedskes/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderAuthority + 1

type initMenu struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return SysBaseMenu{}.TableName()
}

func (i *initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysBaseMenu{},
		&SysBaseMenuParameter{},
		&SysBaseMenuBtn{},
	)
}

func (i *initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysBaseMenu{}) &&
		m.HasTable(&SysBaseMenuParameter{}) &&
		m.HasTable(&SysBaseMenuBtn{})
}

func (i *initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []SysBaseMenu{
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "login", Name: "login", Component: "view/login/index.vue", Sort: 1, Meta: Meta{Title: "登录", Icon: "odometer"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "layout", Name: "layout", Component: "view/layout/index.vue", Sort: 9, Meta: Meta{Title: "布局架子", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 9, Meta: Meta{Title: "关于我们", Icon: "info-filled"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 1, Meta: Meta{Title: "个人信息", Icon: "avatar"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "game", Name: "game", Component: "view/game/index.vue", Sort: 2, Meta: Meta{Title: "比赛页", Icon: "tickets"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "error", Name: "error", Component: "view/error/index.vue", Sort: 5, Meta: Meta{Title: "404", Icon: "upload"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "reload", Name: "reload", Component: "view/error/reload.vue", Sort: 5, Meta: Meta{Title: "reload", Icon: "upload"}},
		{MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: Meta{Title: "个人信息", Icon: "message"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "rank", Name: "rank", Component: "view/rank/index.vue", Sort: 5, Meta: Meta{Title: "排行榜", Icon: "notebook"}},
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "announcement", Name: "announcement", Component: "view/announcement/index.vue", Sort: 5, Meta: Meta{Title: "公告栏", Icon: "notebook"}},

		//管理员菜单：
		{MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "admin", Component: "view/admin/index.vue", Sort: 3, Meta: Meta{Title: "管理员", Icon: "user"}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "gameadmin", Name: "gameadmin", Component: "view/admin/game/index.vue", Sort: 2, Meta: Meta{Title: "比赛管理", Icon: "tickets"}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "user", Name: "user", Component: "view/admin/user/user.vue", Sort: 4, Meta: Meta{Title: "用户管理", Icon: "coordinate"}},
		{MenuLevel: 0, Hidden: false, ParentId: "11", Path: "file", Name: "file", Component: "view/file/index.vue", Sort: 5, Meta: Meta{Title: "文件管理", Icon: "upload"}},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysBaseMenu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "autoPkg").First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
