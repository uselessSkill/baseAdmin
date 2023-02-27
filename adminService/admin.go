package adminService

import (
	"baseAdmin/common"
	"baseAdmin/db/test"
	"baseAdmin/output"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdminUser struct {
}

// 添加
func (*AdminUser) Add(c *gin.Context) {
	var u test.SysAdmin
	if err := c.BindQuery(&u); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if has := test.GetAdmin(0, u.Name); has.Id > 0 {
		output.Json(c, output.AdminExist, output.DefaultData)
		return
	}
	u.SaltPassword()
	u.Ctime = common.GetDateUnix()
	u.Utime = common.GetDateUnix()
	u.Status = 0
	u.LastLogin = ""
	test.SysAdminClient().Create(&u)
	output.Json(c, output.Success, output.DefaultData)
}

// 更新
func (*AdminUser) Update(c *gin.Context) {
	var u test.SysAdmin
	updPwd, _ := strconv.Atoi(c.Query("updPwd"))

	if err := c.BindQuery(&u); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if u.Id < 0 {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if has := test.GetAdmin(0, u.Name); has.Id > 0 {
		output.Json(c, output.AdminExist, output.DefaultData)
		return
	}

	var updField test.SysAdmin

	updField.Name = u.Name
	updField.Id = u.Id
	updField.Nickname = u.Nickname
	updField.Utime = common.GetDateUnix()

	if updPwd == 1 {
		updField.Password = u.Password
		updField.SaltPassword()
	}

	u.Utime = common.GetDateUnix()
	test.SysAdminClient().Where("id =?", updField.Id).Updates(&updField)
	output.Json(c, output.Success, output.DefaultData)
}

// 禁止
func (*AdminUser) Forbid(c *gin.Context) {
	adminId, _ := c.GetQuery("admin_id")

	var admin test.SysAdmin

	admin.Status = 1
	admin.Utime = common.GetDateUnix()

	test.SysAdminClient().Where("id =?", adminId).Updates(&admin)
	output.Json(c, output.Success, output.DefaultData)
}
