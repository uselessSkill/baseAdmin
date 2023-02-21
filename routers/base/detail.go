package base

import (
	"baseAdmin/handle"
	"baseAdmin/output"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NewsRouters(g *gin.Engine) {
	r := g.Group("/base")
	{
		// 稿件详情
		r.GET("/details", func(c *gin.Context) {
			var gids = []string{
				"43fnj457p3p9is90h7ilt7ii5re",
				"44r3hm1ac5n8l7bdqigh3a71m2u",
				"52r3lc1e21l87k8kbg7jfihvlkl",
			}
			newsData := handle.GetDetails(gids)

			c.JSON(http.StatusOK, newsData)
		})
		// 输出
		r.GET("/output", func(c *gin.Context) {
			// 字符串 => 数组
			arr := strings.Split("1,2,3,4", ",")
			output.Json(c, output.Success, arr)
		})

		// 检测输入类型 参数缺失
		r.GET("/checkParams", func(c *gin.Context) {
			handle.Info(c)
		})

		// 读取固定的配置
		r.GET("/config", func(c *gin.Context) {
			handle.GetConfig(c)
		})
		// 读取 需更新的配置信息
		// task 可以对配置进行定期更新
		r.GET("/task", func(c *gin.Context) {
			handle.Task(c)
		})
		// task 的更新可以对此生效
		r.GET("/task/config", func(c *gin.Context) {
			handle.GetTaskConfig(c)
		})
	}
}
