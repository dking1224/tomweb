package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type BaseConfig struct {
	UrlHandler UrlHandler
	Validator  Validator
	ConfigPath string
	DBValue    DBValue
	SqlMethod  SqlMethod
}

type UrlHandler interface {
	ApiHandler(group *gin.RouterGroup)
	CommonHandler(group *gin.RouterGroup)
	AuthHandler(group *gin.RouterGroup)
	UploadHandler(group *gin.RouterGroup)
}

type Validator interface {
	ValidatorParam()
}

type DBValue interface {
	CreateDB() *sqlx.DB
}
