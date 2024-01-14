package main

import (
	"database/sql"
	"fmt"
	"github.com/lindocedskes/core"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/initialize"
	"go.uber.org/zap"
	"os"
)

func main() {
	//router.InitRouter() //初始化路由
	workdirectory, _ := os.Getwd()
	fmt.Printf("当前工作目录：%s \n", workdirectory)

	global.NBUCTF_VP = core.Viper()    //读取配置文件，写入common.NBUCTF_CONFIG
	initialize.OtherInit()             //初始化其他配置: 本地cache及其过期时间
	global.GVA_LOG = core.Zap()        // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG) // 替换zap库为全局变量

	global.NBUCTF_DB = initialize.Gorm()     //连接数据库
	global.NBUCTF_REDIS = initialize.Redis() //连接redis
	fmt.Println("服务端口:", global.NBUCTF_CONFIG.System.Port)

	if global.NBUCTF_DB != nil { //NBUCTF_DB 存配置文件DB数据
		initialize.RegisterTables(global.NBUCTF_DB) //创建或更新数据库表
		initialize.InitSource()                     //表数据初始化 Casbin、Role 和 User
		db, _ := global.NBUCTF_DB.DB()
		defer func(db *sql.DB) { //匿名函数
			err := db.Close()
			if err != nil {
				fmt.Println("db close failed")
			}
		}(db) // db 变量作为参数传入
	}

	//initialize.StartRouter(address) //初始化路由

	core.RunServer() //启动核心服务
}
