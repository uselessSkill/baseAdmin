package controller

import (
	"baseAdmin/common"
	"baseAdmin/model/test"
	"baseAdmin/output"
	"baseAdmin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdminUser struct {
}

type UserBase struct {
	Sa  test.SysAdmin
	RID int `form:"role_id" json:"role_id" binding:"required"`
}

// 添加
func (au *AdminUser) Add(c *gin.Context) {
	var u UserBase
	var rbac service.Rbac
	if err := c.BindQuery(&u); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if has := test.GetAdmin(0, u.Sa.Name); has.Id > 0 {
		output.Json(c, output.AdminExist, output.DefaultData)
		return
	}

	if has := test.GetRole(u.RID, ""); has.Id < 0 {
		output.Json(c, output.RoleNotExist, output.DefaultData)
	}

	u.Sa.SaltPassword()
	u.Sa.Ctime = common.GetDateUnix()
	u.Sa.Utime = common.GetDateUnix()
	u.Sa.Status = 0
	u.Sa.LastLogin = ""
	test.SysAdminClient().Create(&u.Sa)
	rbac.UserBindRole(u.Sa.Id, u.RID)
	output.Json(c, output.Success, output.DefaultData)
}

// 更新
func (au *AdminUser) Update(c *gin.Context) {
	var u UserBase
	updPwd, _ := strconv.Atoi(c.Query("updPwd"))

	if err := c.BindQuery(&u); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if has := test.GetRole(u.RID, ""); has.Id < 0 {
		output.Json(c, output.RoleNotExist, output.DefaultData)
	}

	if u.Sa.Id < 0 {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if has := test.GetAdmin(0, u.Sa.Name); has.Id != u.Sa.Id && has.Id > 0 {
		output.Json(c, output.AdminExist, output.DefaultData)
		return
	}

	var updField test.SysAdmin
	var rbac service.Rbac

	updField.Name = u.Sa.Name
	updField.Id = u.Sa.Id
	updField.Nickname = u.Sa.Nickname
	updField.Utime = common.GetDateUnix()

	if updPwd == 1 {
		updField.Password = u.Sa.Password
		updField.SaltPassword()
	}

	u.Sa.Utime = common.GetDateUnix()
	test.SysAdminClient().Where("id =?", updField.Id).Updates(&updField)
	rbac.UserBindRole(u.Sa.Id, u.RID)
	output.Json(c, output.Success, output.DefaultData)
}

// 禁止
func (au *AdminUser) Forbid(c *gin.Context) {
	adminId, _ := c.GetQuery("admin_id")

	var admin test.SysAdmin

	admin.Status = 1
	admin.Utime = common.GetDateUnix()

	test.SysAdminClient().Where("id =?", adminId).Updates(&admin)
	output.Json(c, output.Success, output.DefaultData)
}

// 单独绑定权限
func (au *AdminUser) SpAuth(c *gin.Context) {
	aId, err := strconv.Atoi(c.Query("auth_id"))
	uId, err1 := strconv.Atoi(c.Query("admin_id"))
	handle := c.Query("handle")

	if err != nil || err1 != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}
	a := test.GetAuth(aId, "")
	if a == nil {
		output.Json(c, output.AuthorNotExist, output.DefaultData)
		return
	}
	if a.MenuType != 11 {
		output.Json(c, output.AuthErrorType, output.DefaultData)
		return
	}

	if a := test.GetAdmin(uId, ""); a == nil {
		output.Json(c, output.AdminNotExist, output.DefaultData)
		return
	}

	var rbac service.Rbac
	if handle == "bind" {
		rbac.UserBindAuth(uId, a.ReqPath)
	}

	if handle == "delete" {
		rbac.UserDeleteAuth(uId, a.ReqPath)
	}

	output.Json(c, output.Success, output.DefaultData)
}

func (au *AdminUser) GetAuth(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("admin_id"))
	RoleID, _ := strconv.Atoi(c.Query("role_id"))
	var rbac service.Rbac

	if RoleID != 0 {
		rbac.UserBindRole(adminID, RoleID)
	}

	res := rbac.GetUserAuth(adminID)
	output.Json(c, output.Success, res)
}

// 如果 给用户附加 权限，其实算是隐式 传递的权限，非直接权限
func (au *AdminUser) Demo(c *gin.Context) {

	//var rbac service.Rbac
	////handle := c.Query("handle")
	//
	//rbac.RoleBindAuth(100, "/rbac/role/del")
	//rbac.RoleBindAuth(100, "/rbac/role/add")
	//rbac.RoleBindAuth(100, "/rbac/role/upd")
	//rbac.UserBindRole(1, 100)
	//rbac.UserBindAuth(1, "/rbac/role/get")
	//
	//
	////res := rbac.CheckUserAuth(1, "/rbac/role/add1")
	//
	//res := rbac.GetUserAuth(1)
	//
	//output.Json(c, output.Success, res)

}
