package adminService

import (
	"baseAdmin/common"
	"baseAdmin/db/test"
	"baseAdmin/output"
)

// 添加角色
func AddRole(r string) (code int) {

	if r := test.GetRole(0, r); r.Id != 0 {
		return output.RoleExist
	}

	var role test.SysRole
	role.Ctime = common.GetDateUnix()
	role.Utime = common.GetDateUnix()
	role.Status = 1
	role.Name = r
	test.SysRoleClient().Create(&role)
	return output.Success
}

// 更新角色
func UpdRole(r *test.SysRole) int {
	if r := test.GetRole(r.Id, ""); r.Id == 0 {
		return output.RoleNotExist
	}

	if r := test.GetRole(0, r.Name); r.Id != 0 {
		return output.RoleExist
	}

	test.SysRoleClient().Model(&r).Where("id = ?", r.Id).Updates(&r)
	return output.Success
}
