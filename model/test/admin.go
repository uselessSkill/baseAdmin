package test

import (
	"baseAdmin/common"
	"baseAdmin/db"
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
	"io"
)

type SysAdmin struct {
	Id        int    `gorm:"id" json:"id" form:"id"`
	Nickname  string `gorm:"nickname" json:"nickname" form:"nickname"`
	Name      string `gorm:"name" json:"name" binding:"required" form:"name"`
	Password  string `gorm:"password" json:"password" form:"password"`
	Salt      string `gorm:"salt" json:"salt"`
	Status    int    `gorm:"status" json:"status"`
	Ctime     string `gorm:"ctime" json:"ctime"`
	Utime     string `gorm:"utime" json:"utime"`
	LastLogin string `gorm:"last_login" json:"last_login"`
}

func SysAdminClient() *gorm.DB {
	return model.MysqlTestClient.Table("sys_admin")
}

func GetAdmin(rId int, r string) *SysAdmin {
	var a *SysAdmin
	if r != "" {
		SysAdminClient().Where("name =?", r).Find(&a)
	}

	if rId != 0 {
		SysAdminClient().Where("id =?", rId).Find(&a)
	}

	return a
}

func (sa *SysAdmin) SaltPassword() {
	MD5 := md5.New()
	sa.Salt = common.RandAllString(8)
	_, _ = io.WriteString(MD5, sa.Password+sa.Salt)
	sa.Password = hex.EncodeToString(MD5.Sum(nil))
}
