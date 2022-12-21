package logger

import (
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// DimeLogPath Dime log file write path
	LogPath      = "/var/log/storyprotocol/storyprotocol.log"
	LogFileMax   = 500     // 500 megabytes
	doLoggerOnce sync.Once // to avoid race condition
	Log          *zap.SugaredLogger
)

// default logger at infolevel
func InitDefaultLogger() *zap.Logger {
	return InitLogger(Levels.Info)
}

func InitLogger(l Level) *zap.Logger {
	level := l.toZapLevel()
	encoder := getEncoder()
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), getLogWriter())
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Log = logger.Sugar()
	zap.ReplaceGlobals(logger)
	return logger
}

func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // e.g. 2021-05-05T04:18:35.334Z
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // INFO, DEBUG
	cfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
	cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(cfg.EncoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	// lumberjack.Logger is safe for concurrent use
	lumberJackLogger := &lumberjack.Logger{
		Filename:   LogPath,
		MaxSize:    LogFileMax,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Default will return Log, will initialize Log once if not yet.
func Default() *zap.SugaredLogger {
	doLoggerOnce.Do(func() {
		if Log == nil {
			InitDefaultLogger()
		}
	})
	return Log
}
