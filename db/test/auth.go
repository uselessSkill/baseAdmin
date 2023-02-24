package test

import (
	"baseAdmin/db"
	"gorm.io/gorm"
)

type SysAuth struct {
	Id       int    `gorm:"id" json:"id" form:"id"`
	Pid      int    `gorm:"pid" json:"pid" form:"pid"`
	Title    string `gorm:"title" json:"title" form:"title"`
	ReqPath  string `gorm:"req_path" json:"req_path" form:"req_path"`
	Icon     string `gorm:"icon" json:"icon" form:"icon"`
	Ctime    string `gorm:"ctime" json:"ctime" form:"ctime"`
	Utime    string `gorm:"utime" json:"utime" form:"utime"`
	Remark   string `gorm:"remark" json:"remark" form:"remark"`
	JumpPath string `gorm:"jump_path" json:"jump_path" form:"jump_path"`
	Sort     int    `gorm:"sort" json:"sort" form:"sort"`
}

func SysAuthClient() *gorm.DB {
	return db.MysqlTestClient.Table("sys_auth")
}

func GetAuth(rId int, r string) *SysAuth {
	var auth *SysAuth
	if r != "" {
		SysAuthClient().Where("title =?", r).Find(&auth)
	}

	if rId != 0 {
		SysAuthClient().Where("id =?", rId).Find(&auth)
	}

	return auth
}
