package adminService

import (
	"baseAdmin/conf"
	"baseAdmin/db"
	"baseAdmin/output"
)

type AuthParams struct {
	Auth   string `form:"auth" json:"auth" binding:"required" validate:"string=1,20"`
	Method string `form:"method" json:"method" binding:"required" validate:"integer=1,10"`
	Role   string `form:"role" json:"role" binding:"required" validate:"string=1,20"`
}

// 添加权限
func AddAuth(a *AuthParams) (code int) {

	db.CreateMysqlTableSheetStruct(conf.LoadConfig.TestDB, "sys_auth")

	//a.Method = strings.ToUpper(a.Method)
	//if ok := rbac.CasEnforcer.AddPolicy(a.Role, a.Auth, a.Method); !ok {
	//	return output.AddAuthorFail
	//}
	return output.Success
}

// 删除权限
func DelAuth(a *AuthParams) (code int) {
	//a.Method = strings.ToUpper(a.Method)
	//if ok := rbac.CasEnforcer.RemovePolicy(a.Role, a.Auth, a.Method); !ok {
	//	return output.RmAuthorFail
	//}
	return output.Success
}
