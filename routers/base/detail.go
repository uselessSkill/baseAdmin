package base

import (
	"baseAdmin/rbac"
	"github.com/gin-gonic/gin"
)

func NewsRouters(g *gin.Engine) {
	r := g.Group("/login")
	{
		// 检测输入类型 参数缺失
		r.GET("/checkParams", func(c *gin.Context) {
			rbac.Info(c)
		})

	}
}
