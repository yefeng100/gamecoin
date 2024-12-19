package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger   *zap.Logger
	logLevel zapcore.Level
	atom     = zap.NewAtomicLevel()

	once sync.Once
)

func InitLog(logFile, level string) {
	once.Do(func() {
		logLevel = getLogLevel(level)
		logger = NewLogger(logFile, level, true).Logger
	})
}

type Logger struct {
	*zap.Logger
	io.Closer
}

func NewLogger(logFile string, level string, defaultAtom bool) Logger {
	f := &lumberjack.Logger{
		Filename:   logFile, // 日志文件路径
		MaxSize:    1024,    // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       // 最多保存days
		Compress:   true,    // 是否压缩 disabled by default
	}
	writer := zapcore.AddSync(f)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		MessageKey:    "M",
		StacktraceKey: "S",

		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	at := atom
	if defaultAtom {
		// use the global log level
		atom.SetLevel(getLogLevel(level))
	} else {
		at := zap.NewAtomicLevel()
		at.SetLevel(getLogLevel(level))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		at, //level,
	)

	caller := zap.AddCaller()
	logger := Logger{
		Logger: zap.New(core, caller, zap.AddCallerSkip(1)),
		Closer: f,
	}
	return logger
}

func funcCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(filepath.Base(caller.FullPath()))
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000 -0700"))
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func Fatal(msg string, fields ...zap.Field) {
	// 根据模式
	logger.Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func FatalT(traceID string, msg string, fields ...zap.Field) {
	// 根据模式
	logger.Fatal(msg, genTraceFields(traceID, fields...)...)
}

func ErrorT(traceID string, msg string, fields ...zap.Field) {
	logger.Error(msg, genTraceFields(traceID, fields...)...)
}

func WarnT(traceID string, msg string, fields ...zap.Field) {
	logger.Warn(msg, genTraceFields(traceID, fields...)...)
}

func DebugT(traceID string, msg string, fields ...zap.Field) {
	logger.Debug(msg, genTraceFields(traceID, fields...)...)
}

func InfoT(traceID string, msg string, fields ...zap.Field) {
	logger.Info(msg, genTraceFields(traceID, fields...)...)
}

// TODO: 性能测试
func genTraceFields(traceID string, fields ...zap.Field) []zap.Field {
	return append([]zap.Field{zap.String("trace_id", traceID)}, fields...)
}

func SetLogLevel(level string) {
	logLevel = getLogLevel(level)
	atom.SetLevel(logLevel)
}

func GetLogger() *zap.Logger {
	return logger
}
