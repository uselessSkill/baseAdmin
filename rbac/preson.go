package rbac

import (
	"baseAdmin/output"

	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
)

type Person struct {
	Name string `form:"name" json:"name" binding:"required" validate:"string=1,5"`
	Sex  int    `form:"sex" json:"sex" binding:"required" validate:"integer=1,2"`
	Age  int    `form:"age" json:"age" binding:"required" validate:"integer=1,20"`
}

// 获取用户信息
func Info(c *gin.Context) {
	var p Person
	if err := c.ShouldBindQuery(&p); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if code, msg := p.CheckAge(c); code > 0 {
		output.JsonSp(c, output.FormatSp(code, msg))
		return
	}
	output.Json(c, output.Success, p)
}

// 检测年龄
func (p *Person) CheckAge(c *gin.Context) (int, string) {
	validator := govalidators.New()
	if err := validator.LazyValidate(p); err != nil {
		return output.VoidAge, err.Error()
	}
	return 0, ""
}
