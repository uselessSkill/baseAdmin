package main

import (
	"baseAdmin/conf"
	"baseAdmin/routers"
)

func main() {
	r := routers.Init()
	_ = r.Run(conf.LoadConfig.Http.Listen)
}
