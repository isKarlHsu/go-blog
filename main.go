package main

import (
	"blog/command"
	"blog/core"
	"blog/global"
	"blog/middleware"
	"blog/router"
	"flag"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化logger配置
	global.Logger = core.InitLogger()
	// 初始化mysql配置
	global.DB = core.InitGorm()
	// 初始化redis
	global.Redis = core.ConnectRedis()

	flag.Parse()
	option := flag.Args()
	optionNum := flag.NArg()
	if optionNum > 0 {
		command.Handle(option)
		return
	}

	r := router.InitRouter()
	r.Use(middleware.Cors())
	global.Logger.Infof("服务运行在: %s", global.Config.System.Addr())
	r.Run(global.Config.System.Addr())
}
