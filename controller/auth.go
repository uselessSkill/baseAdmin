package controller

import (
	"baseAdmin/common"
	"baseAdmin/model/test"
	"baseAdmin/output"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdminAuth struct {
}

// 添加权限
func (ah *AdminAuth) Add(c *gin.Context) {

	var a test.SysAuth
	if err := c.BindQuery(&a); err != nil {
		output.Json(c, output.MissParams, err.Error())
		return
	}

	if h := test.GetAuth(0, a.Title); h == nil || h.Id != 0 {
		output.Json(c, output.AuthorExist, output.DefaultData)
		return
	}
	a.Ctime = common.GetDateUnix()
	a.Utime = common.GetDateUnix()
	test.SysAuthClient().Create(&a)
	output.Json(c, output.Success, a)
}

// 更新权限
func (ah *AdminAuth) Update(c *gin.Context) {

	var a test.SysAuth
	if err := c.BindQuery(&a); err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	if h := test.GetAuth(a.Id, ""); h.Id == 0 {
		output.Json(c, output.AuthorNotExist, a)
		return
	}

	if h := test.GetAuth(0, a.Title); h.Id != 0 {
		output.Json(c, output.AuthorExist, output.DefaultData)
		return
	}

	a.Utime = common.GetDateUnix()
	test.SysAuthClient().Where("id = ?", a.Id).Updates(&a)
	output.Json(c, output.Success, output.DefaultData)
}

// 更新权限
func (ah *AdminAuth) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	var a test.SysAuth
	a.Id = id
	test.SysAuthClient().Where("id = ?", id).Delete(&a)
	output.Json(c, output.Success, output.DefaultData)
}

// 更新权限
func (ah *AdminAuth) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		output.Json(c, output.MissParams, output.DefaultData)
		return
	}

	var a test.SysAuth
	a.Id = id
	test.SysAuthClient().Where("id = ?", id).Take(&a)
	output.Json(c, output.Success, a)
}
