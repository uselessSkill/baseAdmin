package conf

import (
	"github.com/BurntSushi/toml"
)

type loadConfig struct {
	Http   *Http  `toml:"http"`
	TestDB *Mysql `toml:"test"`
}

type Http struct {
	Listen  string `toml:"listen" json:"listen"`
	TimeOut int    `toml:"timeout" json:"timeout"`
}

type Mysql struct {
	Host   string `toml:"host"`
	User   string `toml:"user"`
	Passwd string `toml:"passwd"`
	Port   int    `toml:"port"`
	DBName string `toml:"dbname"`
}

var LoadConfig *loadConfig

func init() {
	LoadConfig = &loadConfig{}
	configFile := "./conf/test/config.toml"
	if _, err := toml.DecodeFile(configFile, &LoadConfig); err != nil {
		panic(err)
	}
}
