package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system/request"
	"gorm.io/gorm"
	"sort"
)

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	Sqlite          = "sqlite"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	InitDataExist   = "\n[%v] --> %v 的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 初始数据成功!\n"
)

const (
	InitOrderSystem   = 10 // 系统初始化顺序初始值
	InitOrderInternal = 1000
	InitOrderExternal = 100000
)

var (
	ErrMissingDBContext        = errors.New("missing db in context")
	ErrMissingDependentContext = errors.New("missing dependent value in context")
	ErrDBTypeMismatch          = errors.New("db type mismatch")
)

// SubInitializer 提供 source/*/init() 使用的接口，每个 initializer 完成一个初始化过程
type SubInitializer interface {
	InitializerName() string                                              // 不一定代表单独一个表，所以改成了更宽泛的语义
	MigrateTable(ctx context.Context) (next context.Context, err error)   // 建表
	InitializeData(ctx context.Context) (next context.Context, err error) // 建数据
	TableCreated(ctx context.Context) bool                                // 表是否已创建
	DataInserted(ctx context.Context) bool                                // 数据是否已插入
}

// TypedDBInitHandler 执行传入的 initializer
type TypedDBInitHandler interface {
	EnsureDB(ctx context.Context, conf *request.InitDB) (context.Context, error) // 建库，失败属于 fatal error，因此让它 panic
	WriteConfig(ctx context.Context) error                                       // 回写配置
	InitTables(ctx context.Context, inits initSlice) error                       // 建表 handler
	InitData(ctx context.Context, inits initSlice) error                         // 建数据 handler
}

// orderedInitializer 组合一个顺序字段，存储和排序所有的初始化过程
type orderedInitializer struct {
	order int
	SubInitializer
}

// initSlice 供 initializer 排序依赖时使用
type initSlice []*orderedInitializer

var (
	initializers initSlice //全局变量，存储所有的 initializer
	cache        map[string]*orderedInitializer
)

// RegisterInit 注册要执行的初始化过程，会在 InitDB() 时调用
func RegisterInit(order int, i SubInitializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
	if cache == nil {
		cache = map[string]*orderedInitializer{}
	}
	name := i.InitializerName()
	if _, existed := cache[name]; existed {
		panic(fmt.Sprintf("Name conflict on %s", name))
	}
	ni := orderedInitializer{order, i}       //创建一个新的 orderedInitializer，存储初始化过程
	initializers = append(initializers, &ni) //将初始化过程加入到全局变量 initializers 中
	cache[name] = &ni                        //将初始化过程加入到全局变量 cache 中
}

/* ---- * service * ---- */

type InitDBService struct{}

// InitDB 创建数据库并初始化 总入口
func (initDBService *InitDBService) InitDB(conf request.InitDB) (err error) {
	ctx := context.TODO()
	fmt.Println(initializers)
	if len(initializers) == 0 {
		return errors.New("无可用初始化过程，请检查初始化是否已执行完成")
	}
	fmt.Println("1_1")
	sort.Sort(&initializers) // 保证有依赖的 initializer 排在后面执行
	// Note: 若 initializer 只有单一依赖，可以写为 B=A+1, C=A+1; 由于 BC 之间没有依赖关系，所以谁先谁后并不影响初始化
	// 若存在多个依赖，可以写为 C=A+B, D=A+B+C, E=A+1;
	// C必然>A|B，因此在AB之后执行，D必然>A|B|C，因此在ABC后执行，而E只依赖A，顺序与CD无关，因此E与CD哪个先执行并不影响
	var initHandler TypedDBInitHandler
	switch conf.DBType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	//case "pgsql":
	//	initHandler = NewPgsqlInitHandler()
	//	ctx = context.WithValue(ctx, "dbtype", "pgsql")
	default:
		initHandler = NewMysqlInitHandler() //创建MysqlInitHandler{} 实例
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	}
	ctx, err = initHandler.EnsureDB(ctx, &conf) // 建库，initHandler是MysqlInitHandler{} 实例，实现TypedDBInitHandler接口
	if err != nil {
		return err
	}

	db := ctx.Value("db").(*gorm.DB)
	global.NBUCTF_DB = db

	if err = initHandler.InitTables(ctx, initializers); err != nil { // 建表
		return err
	} //建表
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	} // 建数据

	if err = initHandler.WriteConfig(ctx); err != nil { // 回写配置
		return err
	} // 回写配置
	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}
	return nil
}

// createDatabase 创建数据库（ EnsureDB() 中调用 ）
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// createTables 创建表（默认 dbInitHandler.initTables 行为）
func createTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.TableCreated(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

/* -- sortable interface -- */

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type InitDB struct {
	DBType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	DBName   string `json:"dbName" binding:"required"`
	DBPath   string `json:"dbPath"`
}
