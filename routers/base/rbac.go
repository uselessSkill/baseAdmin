package base

import (
	"baseAdmin/output"
	"baseAdmin/rbac"
	"github.com/gin-gonic/gin"
)

func RbacRouters(g *gin.Engine) {
	r := g.Group("/rbac")
	{
		// 检测输入类型 参数缺失
		//r.GET("/init", func(c *gin.Context) {
		//	rbac.Init()
		//	c.JSON(http.StatusOK,"success")
		//})
		// 添加权限
		r.GET("/add", func(c *gin.Context) {
			var a *rbac.AuthParams
			if err := c.ShouldBindQuery(&a); err != nil {
				output.Json(c, output.MissParams, output.DefaultData)
			}

			code := rbac.AddAuth(a)
			output.Json(c, code, output.DefaultData)
		})

		// 删除权限
		r.GET("/del", func(c *gin.Context) {
			var a *rbac.AuthParams
			if err := c.ShouldBindQuery(&a); err != nil {
				output.Json(c, output.MissParams, output.DefaultData)
			}
			code := rbac.DelAuth(a)
			output.Json(c, code, output.DefaultData)
		})

		// 删除权限
		r.GET("/get", func(c *gin.Context) {
			r := c.Query("role")
			data := rbac.GetAuth(r)
			output.Json(c, output.Success, data)
		})

	}
}
