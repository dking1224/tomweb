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
		ConfigPath: "E:\\goLandProjects\\tomweb\\conf.json",
	}
	web.StartServer(baseConfig, ":8088")
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
