package adminService

import (
	"baseAdmin/output"
	"baseAdmin/rbac"
	"strings"
)

// 添加权限
func AddRoleAuth(a *AuthParams) (code int) {
	a.Method = strings.ToUpper(a.Method)
	if ok := rbac.CasEnforcer.AddPolicy(a.Role, a.Auth, a.Method); !ok {
		return output.AddAuthorFail
	}
	return output.Success
}

// 删除权限
func DelRoleAuth(a *AuthParams) (code int) {
	a.Method = strings.ToUpper(a.Method)
	if ok := rbac.CasEnforcer.RemovePolicy(a.Role, a.Auth, a.Method); !ok {
		return output.RmAuthorFail
	}
	return output.Success
}
