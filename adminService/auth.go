package adminService

import (
	"baseAdmin/common"
	"baseAdmin/db/test"
	"baseAdmin/output"
	"github.com/gin-gonic/gin"
)

type AdminAuth struct {
}

// 添加权限
func (ah *AdminAuth) AddAuth(c *gin.Context) {

	var a test.SysRole
	if err := c.BindQuery(&a); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if h := test.GetAuth(0, a.Name); h.Id != 0 {
		output.Json(c, output.AuthorExist, output.DefaultData)
		return
	}

	a.Ctime = common.GetDateUnix()
	a.Utime = common.GetDateUnix()
	test.SysRoleClient().Create(&a)

}

func (ah *AdminAuth) UpdAuth() {

}
