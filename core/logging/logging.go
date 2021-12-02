package logging

import (
	"gin-template/core/setting"
	"gin-template/core/util/file"
	"io"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logFile *zap.SugaredLogger
var logConsole *zap.SugaredLogger

var enableFileLog = false

func Setup() {
	enableFileLog = setting.Config.App.Log.File.Enable
	logger, _ := zap.NewDevelopment()
	logConsole = logger.Sugar()

	// 是否打印日志到文件中
	if enableFileLog {
		// 设置gin server的日志输出
		serverFile := CreateLogFile("server.log")
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(os.Stdout, serverFile)

		// 设置application的日志输出
		zapLevel := zapcore.InfoLevel
		level := setting.Config.App.Log.Level
		switch level {
		case "info":
			zapLevel = zapcore.InfoLevel
		case "warn":
			zapLevel = zapcore.WarnLevel
		case "error":
			zapLevel = zapcore.ErrorLevel
		case "fatal":
			zapLevel = zapcore.FatalLevel
		}
		// lumberjack.Logger is already safe for concurrent use, so we don't need to
		// lock it.
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename: path.Join(setting.Config.App.Log.File.Path,
				"application.log"),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		})
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zapLevel,
		)
		logger := zap.New(core)
		logFile = logger.Sugar()

	}

}

func CreateLogFile(fileName string) *os.File {
	logPath := setting.Config.App.Log.File.Path
	if err := file.MkdirIfNotExist(logPath); err != nil {
		Error(err)
	}
	f, err := os.Create(path.Join(logPath, fileName))
	if err != nil {
		Error(err)
	}
	return f
}

func Debug(args ...interface{}) {
	logConsole.Debug(args)
	if enableFileLog {
		logFile.Debug(args)
	}
}

func Debugf(template string, args ...interface{}) {
	logConsole.Debugf(template, args)
	if enableFileLog {
		logFile.Debugf(template, args)
	}
}

func Info(args ...interface{}) {
	logConsole.Info(args)
	if enableFileLog {
		logFile.Info(args)
	}
}

func Infof(template string, args ...interface{}) {
	logConsole.Infof(template, args)
	if enableFileLog {
		logFile.Infof(template, args)
	}
}

func Warn(args ...interface{}) {
	logConsole.Warn(args)
	if enableFileLog {
		logFile.Warn(args)
	}
}

func Warnf(template string, args ...interface{}) {
	logConsole.Warnf(template, args)
	if enableFileLog {
		logFile.Warnf(template, args)
	}
}

func Error(args ...interface{}) {
	logConsole.Error(args)
	if enableFileLog {
		logFile.Error(args)
	}
}

func Errorf(template string, args ...interface{}) {
	logConsole.Errorf(template, args)
	if enableFileLog {
		logFile.Errorf(template, args)
	}
}

func Fatal(args ...interface{}) {
	logConsole.Fatal(args)
	if enableFileLog {
		logFile.Fatal(args)
	}
}

func Fatalf(template string, args ...interface{}) {
	logConsole.Fatalf(template, args)
	if enableFileLog {
		logFile.Fatalf(template, args)
	}
}

func Panic(args ...interface{}) {
	logConsole.Panic(args)
	if enableFileLog {
		logFile.Panic(args)
	}
}

func Panicf(template string, args ...interface{}) {
	logConsole.Panicf(template, args)
	if enableFileLog {
		logFile.Panicf(template, args)
	}
}
