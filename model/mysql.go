package model

import (
	"baseAdmin/conf"
	"fmt"
	"github.com/gohouse/converter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlTestClient *gorm.DB
)

func init() {
	MysqlTestClient = newMysqlCline(FormatDsn(conf.LoadConfig.TestDB))
}

func newMysqlCline(dsn string) *gorm.DB {
	c, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return c
}

func FormatDsn(mysqlStruct *conf.Mysql) string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		mysqlStruct.User,
		mysqlStruct.Passwd,
		mysqlStruct.Host,
		mysqlStruct.Port,
		mysqlStruct.DBName,
	)
	return dsn
}

func CreateMysqlTableSheetStruct(mysqlStruct *conf.Mysql, tableName string) bool {
	err := converter.NewTable2Struct().
		SavePath("./model.go").
		Dsn(FormatDsn(mysqlStruct)).
		TagKey("gorm").
		EnableJsonTag(true).
		Table(tableName).
		Run()

	if err != nil {
		return false
	}
	return true
}
