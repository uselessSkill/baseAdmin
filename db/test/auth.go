package test

import (
	"baseAdmin/db"
	"gorm.io/gorm"
)

type SysAuth struct {
	Id       int    `gorm:"id" json:"id"`
	Pid      int    `gorm:"pid" json:"pid"`
	Name     string `gorm:"name" json:"name"`
	Icon     string `gorm:"icon" json:"icon"`
	Ctime    string `gorm:"ctime" json:"ctime"`
	Utime    string `gorm:"utime" json:"utime"`
	Remark   string `gorm:"remark" json:"remark"`
	JumpPath string `gorm:"jump_path" json:"jump_path"`
	Sort     int    `gorm:"sort" json:"sort"`
}

func SysAuthClient() *gorm.DB {
	return db.MysqlTestClient.Table("sys_auth")
}

func GetAuth(rId int, r string) *SysAuth {
	var auth *SysAuth
	if r != "" {
		SysAuthClient().Where("name =?", r).Find(&auth)
	}

	if rId != 0 {
		SysAuthClient().Where("id =?", rId).Find(&auth)
	}

	return auth
}
