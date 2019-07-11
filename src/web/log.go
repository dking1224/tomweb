package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func InitLog() {
	fmt.Println("log init")
	level, _ := logrus.ParseLevel(Conf.LogConf.LogLevel)
	Log.Level = level
	Log.AddHook(&SourceHook{})
	Log.AddHook(&DefHook{AppName: Conf.ApplicationName})
	maxAge := time.Hour.Nanoseconds() * Conf.LogConf.MaxAge
	rotationTime := time.Hour.Nanoseconds() * Conf.LogConf.RotationTime
	Log.AddHook(ConfigLocalFilesystemLogger(Conf.LogConf.LogPath, Conf.LogConf.LogFileName, time.Duration(maxAge), time.Duration(rotationTime)))
}

//日志的等级
func LogLevel(level logrus.Level) {
	Log.Level = level
}

func Info(tag string, requestID uint64, message interface{}, data map[string]interface{}) {
	Log.WithFields(GetFields(tag, requestID, data)).Info(message)
}

func CError(ctx *gin.Context, err error) {
	value, _ := ctx.Get("requestID")
	Log.WithField("requestID", value).Error(err)
}

func CInfo(ctx *gin.Context, msg string) {
	value, _ := ctx.Get("requestID")
	Log.WithField("requestID", value).Info(msg)
}

func GetFields(tag string, requestID uint64, data map[string]interface{}) logrus.Fields {
	fields := make(map[string]interface{})
	fields["tag"] = tag
	fields["requestID"] = requestID
	for k, v := range data {
		fields[k] = v
	}
	return fields
}
