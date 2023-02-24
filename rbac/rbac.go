package rbac

import (
	"baseAdmin/conf"
	"baseAdmin/db"
	"github.com/casbin/casbin"
	xd "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
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
