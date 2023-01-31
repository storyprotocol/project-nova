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

func (l Level) toZapLevel() (zapcore.Level, error) {
	switch l {
	case Levels.Debug:
		return zap.DebugLevel, nil
	case Levels.Info:
		return zap.InfoLevel, nil
	case Levels.Warn:
		return zap.WarnLevel, nil
	case Levels.Error:
		return zap.ErrorLevel, nil
	case Levels.Dpanic:
		return zap.DPanicLevel, nil
	case Levels.Panic:
		return zap.PanicLevel, nil
	case Levels.Fatal:
		return zap.FatalLevel, nil
	}
	return zap.FatalLevel, fmt.Errorf("unknown logger level: %v", l)
}

var (
	Usage = fmt.Sprintf("log levels: %v|%v|%v|%v", Levels.Info, Levels.Debug, Levels.Warn, Levels.Error)
)
