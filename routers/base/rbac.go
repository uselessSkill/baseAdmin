package base

import (
	"baseAdmin/output"
	"baseAdmin/rbac"
	"github.com/gin-gonic/gin"
)

func RbacRouters(g *gin.Engine) {
	r := g.Group("/rbac")
	{
		// 添加角色
		r.GET("/role/add", func(c *gin.Context) {
			r := c.Query("name")
			code := rbac.AddRole(r)
			output.Json(c, code, output.DefaultData)
		})

		// 添加权限
		r.GET("/auth/add", func(c *gin.Context) {
			var a *rbac.AuthParams
			if err := c.ShouldBindQuery(&a); err != nil {
				output.Json(c, output.MissParams, output.DefaultData)
			}

			code := rbac.AddAuth(a)
			output.Json(c, code, output.DefaultData)
		})

		// 删除权限
		r.GET("/auth/del", func(c *gin.Context) {
			var a *rbac.AuthParams
			if err := c.ShouldBindQuery(&a); err != nil {
				output.Json(c, output.MissParams, output.DefaultData)
			}
			code := rbac.DelAuth(a)
			output.Json(c, code, output.DefaultData)
		})

		// 删除权限
		r.GET("/auth/get", func(c *gin.Context) {
			r := c.Query("role")
			data := rbac.GetRoleAuth(r)
			output.Json(c, output.Success, data)
		})

	}
}
