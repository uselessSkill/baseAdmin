package base

import (
	"baseAdmin/adminService"
	"baseAdmin/db/test"
	"baseAdmin/output"
	"github.com/gin-gonic/gin"
)

func RbacRouters(g *gin.Engine) {
	r := g.Group("/rbac")
	{
		// 添加角色
		r.GET("/role/add", func(c *gin.Context) {
			rName := c.Query("name")
			code := adminService.AddRole(rName)
			output.Json(c, code, output.DefaultData)
		})

		// 更新角色
		r.GET("/role/update", func(c *gin.Context) {

			var roleParams *test.SysRole
			if err := c.ShouldBindQuery(&roleParams); err != nil {
				output.Json(c, output.MissParams, output.DefaultData)
				return
			}

			code := adminService.UpdRole(roleParams)
			output.Json(c, code, output.DefaultData)
		})

		// 添加权限
		r.GET("/auth/add", func(c *gin.Context) {
			var a *adminService.AuthParams
			if err := c.ShouldBindQuery(&a); err != nil {
				output.Json(c, output.MissParams, output.DefaultData)
			}

			code := adminService.AddAuth(a)
			output.Json(c, code, output.DefaultData)
		})

		// 删除权限
		r.GET("/auth/del", func(c *gin.Context) {
			var a *adminService.AuthParams
			if err := c.ShouldBindQuery(&a); err != nil {
				output.Json(c, output.MissParams, output.DefaultData)
			}
			code := adminService.DelAuth(a)
			output.Json(c, code, output.DefaultData)
		})

		//// 获取权限
		//r.GET("/auth/get", func(c *gin.Context) {
		//	r := c.Query("role")
		//	data := adminService.GetRoleAuth(r)
		//	output.Json(c, output.Success, data)
		//})

	}
}
