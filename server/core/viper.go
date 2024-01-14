package core

import (
	"flag" //flag：这是一个用于处理配置文件的库，它提供了读取、解析和监听配置文件变化的功能。
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/core/internal"
	"github.com/lindocedskes/global"
	"github.com/spf13/viper"
	"os"
)

// 读取配置文件
// priority： 1. overrides 2. flags(命令行)3. env. variables 4. config file 5. key/value store 6. defaults
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper() *viper.Viper {
	var config string
	//根据当前的 gin 模式（调试模式、发布模式或测试模式）选择默认的配置文件
	flag.StringVar(&config, "c", "", "choose config file.") //flag：定义一个命令行参数，-c，存储指针、指定名称、默认值和解释说明
	flag.Parse()                                            //解析一个命令行参数
	if config == "" {                                       // 判断命令行参数是否为空
		if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
			switch gin.Mode() {
			case gin.DebugMode:
				config = internal.ConfigDefaultFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.DebugMode, internal.ConfigDefaultFile)
			case gin.ReleaseMode:
				config = internal.ConfigReleaseFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.ReleaseMode, internal.ConfigReleaseFile)
			case gin.TestMode:
				config = internal.ConfigTestFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.TestMode, internal.ConfigTestFile)
			}
		} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
			config = configEnv
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
		}
	} else { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
	}

	// 开始读取 viper 配置
	v := viper.New()
	v.SetConfigFile(config) //设置路径
	v.SetConfigType("yaml")
	err := v.ReadInConfig() //读取
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//fmt.Println(v.Get("mysql.host")) //测试读取配置文件

	v.WatchConfig()                           //监视，如果配置文件被修改，Viper 会自动重新读取
	v.OnConfigChange(func(e fsnotify.Event) { //回调函数,配置文件被修改，对应执行
		fmt.Println("config file changed:", e.Name)
		//Unmarshal():将从配置文件读取的数据解析并填充到 global.NBUCTF_CONFIG 结构体
		if err = v.Unmarshal(&global.NBUCTF_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.NBUCTF_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
