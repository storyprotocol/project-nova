package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level string

var Levels = struct {
	Debug  Level
	Info   Level
	Warn   Level
	Error  Level
	Dpanic Level
	Panic  Level
	Fatal  Level
}{
	Debug:  "debug",
	Info:   "info",
	Warn:   "warn",
	Error:  "error",
	Dpanic: "dpanic",
	Panic:  "panic",
	Fatal:  "fatal",
}

func (l Level) toZapLevel() zapcore.Level {
	switch l {
	case Levels.Debug:
		return zap.DebugLevel
	case Levels.Info:
		return zap.InfoLevel
	case Levels.Warn:
		return zap.WarnLevel
	case Levels.Error:
		return zap.ErrorLevel
	case Levels.Dpanic:
		return zap.DPanicLevel
	case Levels.Panic:
		return zap.PanicLevel
	case Levels.Fatal:
		return zap.FatalLevel
	default:
		panic(fmt.Sprintf("unknown logger level: %v", l))
	}
}

var (
	Usage = fmt.Sprintf("log levels: %v|%v|%v|%v", Levels.Info, Levels.Debug, Levels.Warn, Levels.Error)
)
