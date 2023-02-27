package test

import (
	"baseAdmin/db"
	"errors"
	"gorm.io/gorm"
)

type SysRole struct {
	Id     int    `gorm:"id" json:"id" binding:"required" form:"id" validate:"int=1,20"`
	Name   string `gorm:"name" json:"name" binding:"required" form:"name" validate:"string=1,20"`
	Status int    `gorm:"status" json:"status" binding:"required" form:"status" validate:"int=1,2"`
	Ctime  string `gorm:"ctime" json:"ctime"`
	Utime  string `gorm:"utime" json:"utime"`
}

func SysRoleClient() *gorm.DB {
	return model.MysqlTestClient.Table("sys_role")
}

func GetRole(rId int, r string) *SysRole {
	var role *SysRole
	if r != "" {
		SysRoleClient().Where("name =?", r).Find(&role)
	}

	if rId != 0 {
		SysRoleClient().Where("id =?", rId).Find(&role)
	}

	return role
}

func (u *SysRole) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Name == "admin" {
		return errors.New("admin user not allowed to update")
	}
	return
}
