package conf

import (
	"github.com/BurntSushi/toml"
)

type LoadConfig struct {
	Http struct {
		Listen  string `toml:"listen" json:"listen"`
		TimeOut int    `toml:"timeout" json:"timeout"`
	} `toml:"http"`
}

func Init() *LoadConfig {
	c := &LoadConfig{}
	configFile := "./conf/test/config.toml"
	if _, err := toml.DecodeFile(configFile, &c); err != nil {
		panic(err)
	}
	return c
}
