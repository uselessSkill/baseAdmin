package main

import (
	"baseAdmin/conf"
	"baseAdmin/routers"
)

func main() {
	config := conf.Init()
	r := routers.Init()
	_ = r.Run(config.Http.Listen)
}
