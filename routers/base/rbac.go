package base

import (
	"baseAdmin/adminService"
	"github.com/gin-gonic/gin"
)

func RbacRouters(g *gin.Engine) {
	r := g.Group("/rbac")
	{
		// 添加角色
		r.GET("/role/add", adminService.AddRole)
		// 更新角色
		r.GET("/role/update", adminService.UpdRole)
		// 添加权限
		r.GET("/auth/add", adminService.AddAuth)

	}
}
