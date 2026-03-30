package util

import "github.com/gin-gonic/gin"

type ResponseBuilder struct {
	C *gin.Context
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (rb *ResponseBuilder) Response(httpCode int, msg string, data interface{}) {
	rb.C.JSON(httpCode, Response{
		Msg:  msg,
		Data: data,
	})
}
