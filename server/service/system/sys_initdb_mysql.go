package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/gookit/color"
	"github.com/lindocedskes/config"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system/request"
	"github.com/lindocedskes/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInitHandler struct{}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// 下面4个函数实现 TypedDBInitHandler 接口
// WriteConfig mysql回写配置
func (h MysqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mysql)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.NBUCTF_CONFIG.System.DbType = "mysql"                           // 重置数据库类型
	global.NBUCTF_CONFIG.Mysql = c                                         // 重置数据库配置
	global.NBUCTF_CONFIG.JWT.SigningKey = uuid.Must(uuid.NewV4()).String() // 重置jwt签名密钥
	cs := utils.StructToMap(global.NBUCTF_CONFIG)
	for k, v := range cs { //map
		global.NBUCTF_VP.Set(k, v) // 重置内存中的viper
	}
	return global.NBUCTF_VP.WriteConfig() // 写入配置文件
}

// EnsureDB 创建数据库并初始化 mysql
func (h MysqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" { // 判断数据库类型
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToMysqlConfig()                  // 转换为mysql配置
	next = context.WithValue(ctx, "config", c) // 加入mysql配置
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据

	dsn := conf.MysqlEmptyDsn() // 获取dsn
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		return nil, err
	} // 创建数据库

	var db *gorm.DB
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: true,    // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	//global.NBUCTF_CONFIG.AutoCode.Root, _ = filepath.Abs("..") // 获取项目根目录
	next = context.WithValue(next, "db", db)
	return next, err
}

// 建表
func (h MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

// 初始化数据
func (h MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx) // 创建子上下文
	defer func(c func()) { c() }(cancel)    // 延迟执行cancel，防止子上下文泄露
	for _, init := range inits {            // 遍历初始化器
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, Mysql, init.InitializerName()) //带颜色标记
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Mysql, init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Mysql, init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Mysql)
	return nil
}
