package web

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) logrus.Hook {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"%Y%m%d.log",
		rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),        // 文件最大保存时间
		// rotatelogs.WithRotationCount(365),  // 最多存365个文件
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		fmt.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})
	return lfHook
}

type DefHook struct {
	AppName string
}

func (hook *DefHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = hook.AppName
	return nil
}

func (hook *DefHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

type SourceHook struct {
}

func (hook *SourceHook) Fire(entry *logrus.Entry) error {
	flag := 0
	for i := 6; i <= 11; i++ {
		if pc, file, line, ok := runtime.Caller(i); ok {
			funcName := runtime.FuncForPC(pc).Name()
			if !strings.Contains(funcName, "github.com/Sirupsen/logrus") {
				entry.Data["source"] = strings.Join([]string{path.Base(file), path.Base(funcName), strconv.Itoa(line)}, "|")
				if flag == 1 {
					break
				}
				flag = 1
			}
		}
	}
	return nil
}

func (hook *SourceHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
