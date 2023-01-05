package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func RespError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &RespData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func RespErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &RespData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &RespData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
