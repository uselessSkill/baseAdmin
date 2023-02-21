package base

import (
	"baseAdmin/handle"
	"github.com/gin-gonic/gin"
)

func NewsRouters(g *gin.Engine) {
	r := g.Group("/base")
	{
		// 检测输入类型 参数缺失
		r.GET("/checkParams", func(c *gin.Context) {
			handle.Info(c)
		})

	}
}
