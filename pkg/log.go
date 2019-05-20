package root

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Log 日志对象
type Log struct {
	Logger *logrus.Logger
}

// NewLogStdOut 创建标准输入输出的log对象，用于测试
func NewLogStdOut() *Log {
	l := Log{}
	log := logrus.New()
	l.Logger = log
	return &l
}

type logLevelMap map[string]logrus.Level

func getLogLevelMap() logLevelMap {

	levelMap := logLevelMap{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}

	return levelMap

}

// NewLogrotate 生成带文件切割的日志对象,将info和debug以及error分开
func NewLogrotate(ccnfiger Configer) (*Log, error) {

	var customErr Error

	config, err := ccnfiger.GetConfig()
	if err != nil {
		customErr.Code = "get_config_err"
		customErr.Err = err
		return nil, &customErr
	}

	var baseDir, baseFileName, level string

	if config.Log == nil {
		baseDir = "log"
		baseFileName = "app.log"
		level = "info"
	} else {
		baseDir = config.Log.Dir
		baseFileName = config.Log.FileName
		level = config.Log.Level
	}

	basePath := path.Join(baseDir, baseFileName)

	infoWriter, err := rotatelogs.New(
		basePath+".info.%Y%m%d",
		rotatelogs.WithLinkName(basePath+".info.current"), // 当前日志名称
		rotatelogs.WithRotationTime(24*time.Hour),         // 日志切割时间
		rotatelogs.WithMaxAge(7*24*time.Hour),             // 日志保存时间
	)

	errWriter, err := rotatelogs.New(
		basePath+".error.%Y%m%d",
		rotatelogs.WithLinkName(basePath+".error.current"),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)

	if err != nil {
		return nil, err
	}

	hook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.WarnLevel:  errWriter,
			logrus.ErrorLevel: errWriter,
			logrus.DebugLevel: os.Stdout,
			logrus.PanicLevel: errWriter,
			logrus.FatalLevel: errWriter,
		},
		&logrus.TextFormatter{},
	)

	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	for levelStr, levelValue := range getLogLevelMap() {
		if level == levelStr {
			log.SetLevel(levelValue)
		}
	}

	log.AddHook(hook)

	var l Log
	l.Logger = log
	return &l, nil
}
