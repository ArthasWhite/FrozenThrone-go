package routers

import (
	"github.com/Arthaslixin/FrozenThrone-go/global"
	"github.com/Arthaslixin/FrozenThrone-go/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) // 设置Gin运行等级
	router := gin.Default()
	router.Use(middleware.GinLogger(), middleware.GinRecovery(false))
	InitHelloRoute(router)
	return router
}
