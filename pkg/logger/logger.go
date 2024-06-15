package logger

import (
	"os"

	"go.uber.org/zap"
)

var (
	ZapLogger        *zap.Logger
	ZapSugaredLogger *zap.SugaredLogger
)

// Initialize zap logger and sugared logger
func init() {
	cfg := zap.NewProductionConfig()
	logFile := os.Getenv("APP_LOG_FILE")
	if logFile != "" {
		cfg.OutputPaths = []string{"stderr", logFile}
	}

	ZapLogger = zap.Must(cfg.Build())
	if os.Getenv("APP_ENV") == "development" {
		ZapLogger = zap.Must(zap.NewDevelopment())
	}
	ZapSugaredLogger = ZapLogger.Sugar()
}

func Sync() {
	err := ZapSugaredLogger.Sync()
	if err != nil {
		zap.Error(err)
	}
}

func Info(msg string, keysAndValues ...interface{}) {
	ZapSugaredLogger.Infow(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	ZapSugaredLogger.Debugw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	ZapSugaredLogger.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	ZapSugaredLogger.Errorw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	ZapSugaredLogger.Fatalw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	ZapSugaredLogger.Panicw(msg, keysAndValues...)
}
