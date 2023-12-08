package zap

import (
	"os"
	"path"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	mu        sync.Mutex
	zapLogger *zap.Logger
	cfg       Config
}

func New(cfg Config) *zap.Logger {
	logger := &Logger{
		cfg: cfg,
	}

	loggerConfig := logger.configureLogger()

	loggerConfig.Encoding = "json"

	logger.zapLogger = zap.New(logger.configureCore(loggerConfig))

	return logger.zapLogger
}

func (l *Logger) Named(name string) *zap.Logger {
	return l.zapLogger.Named(name)
}

func (l *Logger) Sync() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.zapLogger.Sync()
}

func (l *Logger) configureLogger() zap.Config {
	loggerConfig := zap.NewDevelopmentConfig()

	if !l.cfg.DevMode {
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.DisableStacktrace = true
	}

	// User friendly field names
	loggerConfig.EncoderConfig.LevelKey = "level"
	loggerConfig.EncoderConfig.NameKey = "module"
	loggerConfig.EncoderConfig.MessageKey = "msg"
	loggerConfig.EncoderConfig.CallerKey = "caller"
	loggerConfig.EncoderConfig.StacktraceKey = "stacktrace"

	return loggerConfig
}

func (l *Logger) configureCore(cfg zap.Config) zapcore.Core {
	fileWS := zapcore.AddSync(&lumberjack.Logger{
		Filename: path.Join(l.cfg.Directory, "logs.jsonl"),
		Compress: true,
	})

	stdoutWS := zapcore.AddSync(os.Stdout)

	mws := zapcore.NewMultiWriteSyncer(stdoutWS, fileWS)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		mws,
		cfg.Level,
	)

	return core
}
