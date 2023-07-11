package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// log file write path
	doLoggerOnce sync.Once // to avoid race condition
	Log          *zap.SugaredLogger
)

// default logger at infolevel
func InitDefaultLogger() (*zap.Logger, error) {
	return InitLogger(Levels.Info)
}

func InitLogger(l Level) (*zap.Logger, error) {
	level, err := l.toZapLevel()
	if err != nil {
		return nil, err
	}
	encoder := getEncoder()
	writeSyncer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Log = logger.Sugar()
	zap.ReplaceGlobals(logger)
	return logger, nil
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

// Default will return Log, will initialize Log once if not yet.
func Default() *zap.SugaredLogger {
	doLoggerOnce.Do(func() {
		if Log == nil {
			_, _ = InitDefaultLogger()
		}
	})
	return Log
}
