package base

import (
	"baseAdmin/adminService"
	"github.com/gin-gonic/gin"
)

func RbacRouters(g *gin.Engine) {
	admin := new(adminService.AllService)
	r := g.Group("/rbac")
	{
		// 添加角色
		r.GET("/role/add", admin.AdminRole.AddRole)
		// 更新角色
		r.GET("/role/update", admin.AdminRole.UpdRole)
		// 添加权限
		r.GET("/auth/add", adminService.AddAuth)

	}
}
