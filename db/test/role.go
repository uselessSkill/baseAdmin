package test

import (
	"baseAdmin/db"
	"gorm.io/gorm"
)

type SysRole struct {
	Id    int    `gorm:"id" json:"id"`
	Name  string `gorm:"name" json:"name"`
	Ctime string `gorm:"ctime" json:"ctime"`
}

var SysRoleClient *gorm.DB

func init() {
	SysRoleClient = db.MysqlTestClient.Table("sys_role")
}
