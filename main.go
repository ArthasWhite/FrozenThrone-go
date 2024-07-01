package main

import (
	"github.com/Arthaslixin/FrozenThrone-go/core"
	"github.com/Arthaslixin/FrozenThrone-go/global"
	"github.com/Arthaslixin/FrozenThrone-go/routers"
)

func main() {
	// 读取配置文件
	global.Config = core.InitConf()
	// log初始化
	global.Logger = core.InitLogger(&global.Config.Logger)
	defer global.Logger.Sync()
	// mysql初始化
	global.DB = core.InitGorm()
	// 路由初始化
	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Logger.Sugar().Infof("项目启动成功,运行在: %s", addr)
	router.Run(addr)
}
