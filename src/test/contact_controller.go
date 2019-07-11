package test

import (
	"github.com/dking1224/tomweb/src/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetById() gin.HandlerFunc {
	return func(context *gin.Context) {
		contactId, ok := web.GetInt(context, "contactId", "0", 10, 64)
		if ok {
			data, error := FindContactById(context, uint64(contactId))
			ok := web.ServiceError(context, error)
			if ok {
				context.JSON(http.StatusOK, web.Success(data, context))
			}
		}
	}
}
