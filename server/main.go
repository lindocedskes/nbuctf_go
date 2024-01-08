package main

import (
	"database/sql"
	"fmt"
	"github.com/lindocedskes/common"
	"github.com/lindocedskes/initialize"
	"os"
)

func main() {
	//routes.InitRouter() //初始化路由
	workdirectory, _ := os.Getwd()
	fmt.Printf("当前工作目录：%s \n", workdirectory)

	common.NBUCTF_VP = initialize.Viper() //读取配置文件，写入common.NBUCTF_CONFIG
	common.NBUCTF_DB = initialize.Gorm()  //连接数据库
	common.NBUCTF_REDIS = initialize.Redis()
	fmt.Println("服务端口:", common.NBUCTF_CONFIG.System.Port)

	if common.NBUCTF_DB != nil { //NBUCTF_DB 存配置文件DB数据
		initialize.RegisterTables(common.NBUCTF_DB) //创建或更新数据库表
		initialize.InitSource()                     //表数据初始化 Casbin、Role 和 User
		db, _ := common.NBUCTF_DB.DB()
		defer func(db *sql.DB) { //匿名函数
			err := db.Close()
			if err != nil {
				fmt.Println("db close failed")
			}
		}(db) // db 变量作为参数传入
	}

	//TODO core.RunServer() //启动核心服务
}
