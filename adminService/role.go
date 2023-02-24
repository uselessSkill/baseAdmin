package adminService

import (
	"baseAdmin/common"
	"baseAdmin/db/test"
	"baseAdmin/output"
	"github.com/gin-gonic/gin"
)

type AdminRole struct {
}

// 添加角色
func (rl *AdminRole) Add(c *gin.Context) {

	r := c.Query("name")

	if r := test.GetRole(0, r); r.Id != 0 {
		output.Json(c, output.RoleExist, output.DefaultData)
		return
	}

	var role test.SysRole
	role.Ctime = common.GetDateUnix()
	role.Utime = common.GetDateUnix()
	role.Status = 1
	role.Name = r
	test.SysRoleClient().Create(&role)
	output.Json(c, output.Success, output.DefaultData)
}

// 更新角色
func (rl *AdminRole) Update(c *gin.Context) {

	var r test.SysRole
	if err := c.BindQuery(&r); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if r := test.GetRole(r.Id, ""); r.Id == 0 {
		output.Json(c, output.RoleNotExist, output.DefaultData)
		return
	}

	if r := test.GetRole(0, r.Name); r.Id != 0 {
		output.Json(c, output.RoleExist, output.DefaultData)
		return
	}

	test.SysRoleClient().Model(&r).Where("id = ?", r.Id).Updates(&r)
	output.Json(c, output.Success, output.DefaultData)
}
