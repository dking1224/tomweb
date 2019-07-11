package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//执行service方法处理错误信息
func ServiceError(context *gin.Context, err error) bool {
	if err != nil {
		context.JSON(http.StatusOK, Fail(context, err.Error()))
		return false
	}
	return true
}

//同时支持get post
func GetAndPost(group *gin.RouterGroup, path string, handlerFunc gin.HandlerFunc) {
	group.POST(path, handlerFunc)
	group.GET(path, handlerFunc)
}
