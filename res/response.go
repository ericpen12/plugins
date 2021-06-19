package res

import (
	"github.com/gin-gonic/gin"
	e "github.com/quzhen12/plugins/error"
)

var (
	defaultErrCode = 1
)

type message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Json(c *gin.Context, data interface{}, err error) {
	m := msg(data, err)
	c.JSON(e.HttpStatusCode(m.Code), m)
}

func msg(data interface{}, err interface{}) message {
	var m message
	switch v := err.(type) {
	case e.Err:
		m.Code = v.Code
		m.Msg = v.Msg
		return m
	case error:
		m.Msg = v.Error()
		m.Code = defaultErrCode
		return m
	}
	switch v := data.(type) {
	case string:
		m.Msg = v
	default:
		m.Data = data
	}
	return m
}
