package rbac

import (
	"baseAdmin/output"
	"github.com/casbin/casbin"
	xd "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var (
	CasEnforcer *casbin.Enforcer
)

func init() {
	// 使用自己定义rbac_db
	// 最后的一个参数咱们写true ，否则默认为false,使用缺省的数据库名casbin,不存在则创建
	a := xd.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/mycasbin?charset=utf8", true)
	CasEnforcer = casbin.NewEnforcer("./conf/auth_model.conf", a)
	//从DB中 load 策略
	_ = CasEnforcer.LoadPolicy()

}

type AuthParams struct {
	Auth   string `form:"auth" json:"auth" binding:"required" validate:"string=1,20"`
	Method string `form:"method" json:"sex" binding:"required" validate:"integer=1,10"`
	Role   string `form:"role" json:"role" binding:"required" validate:"integer=1,20"`
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

func GetAuth(r string) interface{} {
	return CasEnforcer.GetFilteredPolicy(0, r)
}

// myAuth 拦截器
func myAuth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		// 获取方法
		act := c.Request.Method
		sub := "root"

		// 判断策略是否已经存在了
		if ok := e.Enforce(sub, obj, act); ok {
			log.Println("Check successfully")
			c.Next()
		} else {
			log.Println("sorry , Check failed")
			c.Abort()
		}
	}
}
