package handle

import (
	"baseAdmin/conf"
	"baseAdmin/output"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	cf := conf.Init()
	output.Json(c, output.Success, cf)
}
