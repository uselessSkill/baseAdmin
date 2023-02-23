package rbac

import (
	"baseAdmin/common"
	"baseAdmin/conf"
	"baseAdmin/db"
	"baseAdmin/db/test"
	"baseAdmin/output"
	"github.com/casbin/casbin"
	xd "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	CasEnforcer *casbin.Enforcer
)

func init() {
	conn := db.FormatDsn(conf.LoadConfig.TestDB)
	// 使用自己定义rbac_db
	// 最后的一个参数咱们写true ，否则默认为false,使用缺省的数据库名casbin,不存在则创建
	a := xd.NewAdapter("mysql", conn+"?charset=utf8", true)
	CasEnforcer = casbin.NewEnforcer("./conf/auth_model.conf", a)
	//从DB中 load 策略
	_ = CasEnforcer.LoadPolicy()

}

type AuthParams struct {
	Auth   string `form:"auth" json:"auth" binding:"required" validate:"string=1,20"`
	Method string `form:"method" json:"method" binding:"required" validate:"integer=1,10"`
	Role   string `form:"role" json:"role" binding:"required" validate:"string=1,20"`
}

// 添加角色
func AddRole(r string) (code int) {
	var role test.SysRole
	test.SysRoleClient.Where("name =?", r).Find(&role)

	if role.Id != 0 {
		return output.RoleExist
	}

	role.Name = r
	role.Ctime = common.GetDateUnix()
	test.SysRoleClient.Table("sys_role").Create(&role)
	return output.Success
}

// 添加权限
func AddAuth(a *AuthParams) (code int) {
	a.Method = strings.ToUpper(a.Method)
	if ok := CasEnforcer.AddPolicy(a.Role, a.Auth, a.Method); !ok {
		return output.AddAuthorFail
	}
	return output.Success
}

// 删除权限
func DelAuth(a *AuthParams) (code int) {
	a.Method = strings.ToUpper(a.Method)
	if ok := CasEnforcer.RemovePolicy(a.Role, a.Auth, a.Method); !ok {
		return output.RmAuthorFail
	}
	return output.Success
}

// 获取角色所属权限
func GetRoleAuth(r string) interface{} {
	return CasEnforcer.GetFilteredPolicy(0, r)
}
