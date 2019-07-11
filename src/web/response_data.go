package web

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type ResponseData struct {
	Code      int8        `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
	RequestId uint64      `json:"requestId"`
}

func InitResponse(code int8, msg string, data interface{}) *ResponseData {
	response := &ResponseData{Code: code, Msg: msg, Data: data}
	return response
}

func SuccessResponse() *ResponseData {
	response := &ResponseData{Code: 0, Msg: "success"}
	return response
}

func SuccessResponseData(data interface{}) *ResponseData {
	response := SuccessResponse()
	if data != nil {
		if reflect.TypeOf(data).Kind() == reflect.Ptr {
			if !reflect.ValueOf(data).IsNil() {
				response.Data = data
			}
		} else {
			response.Data = data
		}
	}
	return response
}

func Success(data interface{}, context *gin.Context) *ResponseData {
	response := SuccessResponseData(data)
	if v, ok := context.Get("requestID"); ok {
		response.RequestId = v.(uint64)
	}
	return response
}

func FailResponse() *ResponseData {
	return FailMsgResponse("fail")
}

func FailMsgResponse(msg string) *ResponseData {
	response := &ResponseData{Code: -1, Msg: msg}
	return response
}

func Fail(ctx *gin.Context, msg string) *ResponseData {
	response := &ResponseData{Code: -1, Msg: msg}
	if v, ok := ctx.Get("requestID"); ok {
		response.RequestId = v.(uint64)
	}
	return response
}
