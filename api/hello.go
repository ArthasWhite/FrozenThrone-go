package api

import (
	"github.com/Arthaslixin/FrozenThrone-go/structure"
	"github.com/gin-gonic/gin"
)

type HelloApi struct {
}

func (h *HelloApi) Hello(c *gin.Context) {
	res := structure.HelloResponse{
		Msg: "hello, world!",
	}
	c.JSON(200, res)
}
