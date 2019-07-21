package test

import (
	"fmt"
	"github.com/dking1224/tomweb/src/web"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestStartServer(t *testing.T) {
	fmt.Println("start")
	baseConfig := &web.BaseConfig{
		UrlHandler: &UrlHandlerTest{},
		DBValue:    &MysqlDBValue{},
		ConfigPath: "E:\\GoProjects\\tomweb\\conf.json",
	}
	fmt.Println(CheckError())
	fmt.Println(BusinessError())
	SelectError(CheckError())
	SelectError(BusinessError())
	web.StartServer(baseConfig, ":8088")
}

func CheckError() error {
	return &web.CheckError{ErrorMsg: "msg"}
}

func BusinessError() error {
	return &web.BusinessError{ErrorMsg: "business"}
}

func SelectError(err error) {
	switch err.(type) {
	case *web.CheckError:
		fmt.Println("check")
	case *web.BusinessError:
		fmt.Println("business")
	default:
		fmt.Println("unknown error")
	}
}

type UrlHandlerTest struct {
}

func (url UrlHandlerTest) ApiHandler(group *gin.RouterGroup) {
	web.GetAndPost(group, "/getById", GetById())
}
func (url UrlHandlerTest) CommonHandler(group *gin.RouterGroup) {
	fmt.Println(web.Conf.Property["test"])
	web.GetAndPost(group, "/health", Health())
}
func (url UrlHandlerTest) AuthHandler(group *gin.RouterGroup) {

}
func (url UrlHandlerTest) UploadHandler(group *gin.RouterGroup) {

}

type MysqlDBValue struct {
}

func (mysqlDB MysqlDBValue) CreateDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", web.Conf.DbConf.DbSource)
	if err != nil {
		web.Log.Error("db connect error", err)
	}
	if web.Conf.DbConf.MaxOpen != 0 {
		db.SetMaxOpenConns(web.Conf.DbConf.MaxOpen)
	}
	if web.Conf.DbConf.MaxIdle != 0 {
		db.SetMaxIdleConns(web.Conf.DbConf.MaxIdle)
	}
	return db
}
