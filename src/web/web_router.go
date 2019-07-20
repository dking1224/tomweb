package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"html"
	"net/http"
	"strconv"
	"time"
)

var Conf *Config
var Log = logrus.New()
var DB *sqlx.DB

//启动服务器
func StartServer(baseConfig *BaseConfig, addr ...string) {
	//配置初始化
	Conf = NewConfig(baseConfig.ConfigPath)
	//日志初始化
	InitLog()
	//sql文件处理
	if Conf.SqlTemplate != "" {
		InitTemplate(Conf.SqlTemplate, baseConfig.SqlMethod)
	}
	DB = baseConfig.DBValue.CreateDB()
	router := gin.Default()
	//参数校验
	if baseConfig.Validator != nil {
		baseConfig.Validator.ValidatorParam()
	}
	//xxs
	router.Use(xxs())
	//请求统计
	router.Use(access())
	//静态文件
	staticHtml(router)
	//错误处理
	errorHandler(router)
	//接口处理
	apiHandler(router, baseConfig.UrlHandler)
	router.Run(addr...)
}

//静态页面
func staticHtml(router *gin.Engine) {
	if config.StaticSource != "" {
		router.LoadHTMLGlob(config.StaticSource)
		router.GET("/", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", nil)
		})
	}
}

//错误处理
func errorHandler(router *gin.Engine) {
	router.NoRoute(func(context *gin.Context) {
		if Conf.StaticSource != "" {
			context.HTML(http.StatusBadRequest, "404.html", nil)
		} else {
			context.JSON(http.StatusOK, FailMsgResponse("404 no route"))
		}
	})

	router.NoMethod(func(context *gin.Context) {
		if Conf.StaticSource != "" {
			context.HTML(http.StatusForbidden, "403.html", nil)
		} else {
			context.JSON(http.StatusOK, FailMsgResponse("405 no method"))
		}
	})
}

//接口api
func apiHandler(router *gin.Engine, url UrlHandler) {
	//无需登陆api
	commonGroup := router.Group("/common")
	//需要登陆api
	apiGroup := router.Group("/api")
	//上传api
	uploadGroup := router.Group("/upload")
	//需要认证api
	authGroup := router.Group("/auth")
	url.ApiHandler(apiGroup)
	url.CommonHandler(commonGroup)
	url.UploadHandler(uploadGroup)
	url.AuthHandler(authGroup)
}

//访问打印访问时间与返回头信息requestID
func access() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestID := GetUUID()
		startTime := time.Now()
		context.Set("requestID", requestID)
		context.Writer.Header().Set("requestID", strconv.FormatUint(requestID, 10))
		Info("access", requestID, "access start", map[string]interface{}{"path": context.Request.URL.Path})
		context.Next()
		total := time.Since(startTime)
		Info("access", requestID, "access end", map[string]interface{}{"costTime": total})
	}
}

//xxs注入
func xxs() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.DefaultPostForm("xxs", "0")
		if form := context.Request.PostForm; form != nil {
			for k, v := range form {
				str := v
				for i, item := range v {
					str[i] = html.EscapeString(item)
				}
				form[k] = str
			}
		}
		if form := context.Request.Form; form != nil {
			for k, v := range form {
				str := v
				for i, item := range v {
					str[i] = html.EscapeString(item)
				}
				form[k] = str
			}
		}
	}
}
