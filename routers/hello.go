package routers

import (
	"github.com/Arthaslixin/FrozenThrone-go/api"
	"github.com/gin-gonic/gin"
)

func InitHelloRoute(r *gin.Engine) {
	helloApi := &api.HelloApi{}
	hello := r.Group("hello")
	{
		hello.GET("hello", helloApi.Hello)
	}
}
