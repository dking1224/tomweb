package test

import (
	"github.com/dking1224/tomweb/src/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health() gin.HandlerFunc {
	return func(context *gin.Context) {
		web.CInfo(context, "health")
		context.JSON(http.StatusOK, web.Success(nil, context))
	}
}
