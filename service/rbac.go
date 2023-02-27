package service

import (
	"baseAdmin/conf"
	"baseAdmin/model"
	"baseAdmin/model/test"
	"baseAdmin/output"
	"github.com/casbin/casbin"
	xd "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Rbac struct {
}

var (
	CasEnforcer *casbin.Enforcer
)

func init() {
	conn := model.FormatDsn(conf.LoadConfig.TestDB)
	// 使用自己定义rbac_db
	// 最后的一个参数咱们写true ，否则默认为false,使用缺省的数据库名casbin,不存在则创建
	a := xd.NewAdapter("mysql", conn+"?charset=utf8", true)
	CasEnforcer = casbin.NewEnforcer("./conf/auth_model.conf", a)
	//从DB中 load 策略
	_ = CasEnforcer.LoadPolicy()
}

// 绑定用户
func (m *Rbac) bindRole(c *gin.Context, rId int) int {

	if has := test.GetRole(rId, ""); has.Id < 0 {
		output.Json(c, output.RoleNotExist, output.DefaultData)
	}

	return output.Success
}
