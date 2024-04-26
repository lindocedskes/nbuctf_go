package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/lindocedskes/core"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/initialize"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	//router.InitRouter() //初始化路由
	workdirectory, _ := os.Getwd()
	fmt.Printf("当前工作目录：%s \n", workdirectory)

	global.NBUCTF_VP = core.Viper()    //读取配置文件，写入common.NBUCTF_CONFIG
	initialize.OtherInit()             //初始化其他配置: 本地cache及其过期时间
	global.GVA_LOG = core.Zap()        // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG) // 替换zap库为全局变量

	global.NBUCTF_DB = initialize.Gorm() //连接数据库
	fmt.Println("服务端口:", global.NBUCTF_CONFIG.System.Port)

	if global.NBUCTF_DB != nil { //NBUCTF_DB 存配置文件DB数据
		initialize.RegisterTables(global.NBUCTF_DB) //创建或更新数据库表（不加数据）
		//initialize.InitSource()                     //表数据初始化 Casbin、Role 和 User （废弃），改用路由执行初始化器
		db, _ := global.NBUCTF_DB.DB()
		defer func(db *sql.DB) { //匿名函数
			err := db.Close()
			if err != nil {
				fmt.Println("db close failed")
			}
		}(db) // db 变量作为参数传入
	}

	var err error
	global.NBUCTF_K8S, err = core.K8s()
	if err != nil {
		fmt.Printf("Failed to initialize Kubernetes client: %v\n", err)
		return
	}
	// 测试 Kubernetes 客户端 是否通信成功
	ctx := context.TODO()
	namespaces, err := global.NBUCTF_K8S.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Failed to list namespaces: %v\n", err)
		return
	}

	for _, namespace := range namespaces.Items {
		fmt.Printf("Namespace: %s\n", namespace.Name)
	}

	core.RunServer() //启动核心服务
}
