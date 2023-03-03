package controller

import (
	"baseAdmin/common"
	"baseAdmin/model/test"
	"baseAdmin/output"
	"baseAdmin/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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

// 操作权限
func (rl *AdminRole) WithAuth(c *gin.Context) {
	rId, _ := strconv.Atoi(c.Query("role_id"))
	authIds := c.Query("auth_ids")
	handle := c.Query("handle")

	if has := test.GetRole(rId, ""); has == nil {
		output.Json(c, output.RoleNotExist, output.DefaultData)
		return
	}

	authMap := strings.Split(authIds, ",")

	var authPathMap []string
	for _, v := range authMap {
		authId, _ := strconv.Atoi(v)
		if has := test.GetAuth(authId, ""); has != nil {
			authPathMap = append(authPathMap, has.ReqPath)
		}
	}

	var rs service.Rbac
	for _, v := range authPathMap {
		if handle == "bind" {
			rs.RoleBindAuth(rId, v)
		} else {
			rs.RoleDeleteAuth(rId, v)
		}
	}

	output.Json(c, output.Success, output.DefaultData)
}
