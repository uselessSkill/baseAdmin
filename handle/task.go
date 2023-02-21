package handle

import (
	"baseAdmin/output"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var config int

func Task(c context.Context) {
	read()
}

func GetTaskConfig(c *gin.Context) {
	// 修改是有效的
	output.Json(c, output.Success, config)
}

func read() {
	//创建一个周期性的定时器
	// 每隔3s发送一次
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for {
			//从定时器中获取数据
			t := <-ticker.C
			config = config + 1
			fmt.Println("当前时间为:", t)
		}
	}()

	for {
		time.Sleep(time.Second * 1)
	}
}
