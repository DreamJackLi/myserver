package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugarLogger *zap.SugaredLogger
	zapConfig   *zap.Config
)

func InitLog() {
	core := zapcore.NewCore(getEncoder(), os.Stdout, zapcore.DebugLevel)
	logger := zap.New(core)
	sugarLogger = logger.Sugar()
	sugarLogger = sugarLogger.With(zap.String("hello", "11111"))
	sugarLogger.With(zap.String("hello1", "11111"))
	sugarLogger.With(zap.String("hello2", "11111"))
	sugarLogger.Info("hello world")

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
