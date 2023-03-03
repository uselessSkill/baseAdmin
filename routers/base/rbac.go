package base

import (
	"baseAdmin/controller"
	"github.com/gin-gonic/gin"
)

func RbacRouters(g *gin.Engine) {
	admin := new(controller.AllService)
	r := g.Group("/rbac")
	{
		// 角色
		r.GET("/role/add", admin.AdminRole.Add)           // 添加角色
		r.GET("/role/update", admin.AdminRole.Update)     // 更新角色
		r.GET("/role/withAuth", admin.AdminRole.WithAuth) // 绑定权限

		// 权限
		r.GET("/auth/add", admin.AdminAuth.Add)       // 添加权限
		r.GET("/auth/update", admin.AdminAuth.Update) // 更新权限
		r.GET("/auth/delete", admin.AdminAuth.Delete) // 删除权限
		r.GET("/auth/get", admin.AdminAuth.Get)       // 获取权限

		// 用户
		r.GET("/admin/add", admin.AdminUser.Add)         // 添加用户
		r.GET("/admin/update", admin.AdminUser.Update)   // 添加用户
		r.GET("/admin/forbid", admin.AdminUser.Forbid)   // 禁用用户
		r.GET("/admin/bindAuth", admin.AdminUser.SpAuth) // 单独处理权限
		r.GET("/admin/getAuth", admin.AdminUser.GetAuth) // 获取用户所有权限
		r.GET("/admin/demo", admin.AdminUser.Demo)       // 获取用户所有权限
	}
}
