package base

import (
	"baseAdmin/adminService"
	"github.com/gin-gonic/gin"
)

func RbacRouters(g *gin.Engine) {
	admin := new(adminService.AllService)
	r := g.Group("/rbac")
	{
		r.GET("/role/add", admin.AdminRole.Add)       // 添加角色
		r.GET("/role/update", admin.AdminRole.Update) // 更新角色
		r.GET("/auth/add", admin.AdminAuth.Add)       // 添加权限
		r.GET("/auth/update", admin.AdminAuth.Update) // 更新权限
		r.GET("/auth/delete", admin.AdminAuth.Delete) // 更新权限
		r.GET("/auth/get", admin.AdminAuth.Get)       // 更新权限
	}
}
