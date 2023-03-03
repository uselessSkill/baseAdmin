package service

import (
	"baseAdmin/conf"
	bm "baseAdmin/model"
	"github.com/casbin/casbin/v2"
	_ "github.com/casbin/casbin/v2/model"
	xd "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type Rbac struct {
}

var (
	CasEnforcer *casbin.Enforcer
)

func init() {
	conn := bm.FormatDsn(conf.LoadConfig.TestDB)
	// 使用自己定义rbac_db
	// 最后的一个参数咱们写true ，否则默认为false,使用缺省的数据库名casbin,不存在则创建
	a, _ := xd.NewAdapter("mysql", conn+"?charset=utf8", true)

	Cef, err := casbin.NewEnforcer("conf/auth_model.conf", a)
	//从DB中 load 策略

	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	_ = Cef.LoadPolicy()
	CasEnforcer = Cef
}

// 绑定用户 g
// 用户 角色 1v1
func (m *Rbac) UserBindRole(uId int, rId int) {
	uIdStr := strconv.Itoa(uId)
	rIdStr := strconv.Itoa(rId)

	_, _ = CasEnforcer.DeleteRolesForUser("u:" + uIdStr)
	_, _ = CasEnforcer.AddRoleForUser("u:"+uIdStr, "r:"+rIdStr)
}

// 角色绑定 权限
func (m *Rbac) RoleBindAuth(id int, other ...string) {
	_, err := CasEnforcer.AddPermissionForUser("r:"+strconv.Itoa(id), other...)
	if err != nil {
		log.Println(err)
	}
}

// 角色 删除 权限
func (m *Rbac) RoleDeleteAuth(id int, other ...string) {
	_, _ = CasEnforcer.DeletePermissionForUser("r:"+strconv.Itoa(id), other...)
}

// 用户 删除 权限
func (m *Rbac) UserDeleteAuth(id int, other ...string) {
	_, _ = CasEnforcer.DeletePermissionForUser("u:"+strconv.Itoa(id), other...)
}

//用户 绑定 权限
func (m *Rbac) UserBindAuth(uId int, other ...string) {
	_, _ = CasEnforcer.AddPermissionForUser("u:"+strconv.Itoa(uId), other...)
}

// 获取用户权限
func (m *Rbac) GetUserAuth(uId int) [][]string {
	res, _ := CasEnforcer.GetImplicitResourcesForUser("u:" + strconv.Itoa(uId))
	return res
}

//获取用户的角色
func (m *Rbac) GetUserRole(UID int) []string {
	res, _ := CasEnforcer.GetRolesForUser("u:" + strconv.Itoa(UID))
	return res
}

func (m *Rbac) CheckUserAuth(uId int, auth string) bool {
	idStr := strconv.Itoa(uId)
	checkRes := CasEnforcer.HasPermissionForUser("u:"+idStr, auth)

	if checkRes == false {
		implicitPermissions, _ := CasEnforcer.GetImplicitPermissionsForUser("u:" + idStr)
		for _, v := range implicitPermissions {
			if v[1] == auth {
				return true
			}
		}
	}

	return checkRes
}
