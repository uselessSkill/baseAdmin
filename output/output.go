package output

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success    = 0
	MissParams = 1000 // 缺少参数
	VoidPrams  = 1001 // 无效的参数
	VoidAge    = 1002 // 无效的年龄
	TimeOut    = 1003
	UnKnow     = 9999
)

var statusText = map[int]string{
	Success:    "ok",
	MissParams: "miss params",
	UnKnow:     "未知错误",
	TimeOut:    "访问超时",
}

var DefaultData = make([]int, 0)

type Resp struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

func Json(c *gin.Context, errno int, data interface{}) {
	c.JSON(http.StatusOK, FormatNor(errno, data))
}

func JsonSp(c *gin.Context, resp *Resp) {
	c.JSON(http.StatusOK, resp)
}

// 常规输出
func FormatNor(errno int, data interface{}) *Resp {
	msg, err := statusText[errno]
	if err != true {
		msg = statusText[UnKnow]
	}
	return &Resp{
		Code:    errno,
		Message: msg,
		Data:    data,
	}
}

// 自定义返回错误
func FormatSp(errno int, msg string) *Resp {
	return &Resp{
		Code:    errno,
		Message: msg,
		Data:    DefaultData,
	}
}

//func (op *Resp) SetMsg(msg string) {
//	op.Message = msg
//}
//
//func (op *Resp) SetMsgF(msg string, v ...interface{}) *Resp {
//	op.Message = fmt.Sprintf(msg, v...)
//	return op
//}
