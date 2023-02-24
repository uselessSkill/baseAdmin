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

	AddAuthorFail = 20001 // 添加权限失败
	RmAuthorFail  = 20002 // 删除权限失败

	AddRoleFail  = 20003 //添加角色失败
	RoleExist    = 20004 //添加角色失败
	RoleNotExist = 20006 //添加角色失败
	RmRoleFail   = 20005 // 删除角色势失败

	UnKnow = 9999
)

var statusText = map[int]string{
	Success:       "ok",
	MissParams:    "miss params",
	UnKnow:        "未知错误",
	TimeOut:       "访问超时",
	AddAuthorFail: "添加权限失败",
	RmAuthorFail:  "删除权限失败",
	AddRoleFail:   "添加角色失败",
	RoleExist:     "角色已存在",
	RoleNotExist:  "角色不存在",
	RmRoleFail:    "删除角色势失败",
}

var DefaultData = make([]int, 0)

type Resp struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

func Json(c *gin.Context, errno int, data interface{}) {
	c.JSON(http.StatusOK, FormatNor(errno, data))
	c.Abort()
}

func JsonSp(c *gin.Context, resp *Resp) {
	c.JSON(http.StatusOK, resp)
	c.Abort()
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
