package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//解析获取json
func GetJSON(context *gin.Context, data interface{}) bool {
	error := context.BindJSON(data)
	if error != nil {
		CError(context, error)
		context.JSON(http.StatusOK, Fail(context, error.Error()))
		return false
	}
	return true
}

//获取int
func GetInt(context *gin.Context, key string, defaultValue string, base int, bitSize int) (int64, bool) {
	contactId := context.DefaultPostForm(key, defaultValue)
	value, err := strconv.ParseInt(contactId, base, bitSize)
	if err != nil {
		CError(context, err)
		context.JSON(http.StatusOK, Fail(context, err.Error()))
		return 0, false
	}
	return value, true
}
