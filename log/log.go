package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

const (
	appModeDebug = "debug"
	logLevelInfo = "info"
)

var (
	logLevel   = logLevelInfo
	filename   = time.Now().Format("2006-06-01-15") + ".log"
	maxSize    = 100
	maxBackups = 7
	maxAge     = 30
	compress   = true
	appMode    = appModeDebug
)

func loadConfig() {
	if viper.GetString("log.level") != "" {
		logLevel = viper.GetString("log.level")
	}
	if viper.GetString("log.filename") != "" {
		filename = viper.GetString("log.filename")
	}
	if viper.GetInt("log.max_size") != 0 {
		maxSize = viper.GetInt("log.max_size")
	}
	if viper.GetInt("log.max_backups") != 0 {
		maxBackups = viper.GetInt("log.max_backups")
	}
	if viper.GetInt("log.max_age") != 0 {
		maxAge = viper.GetInt("log.max_age")
	}
	if viper.GetBool("log.compress") {
		compress = viper.GetBool("log.compress")
	}
	if viper.GetString("app.mode") != "" {
		appMode = viper.GetString("app.mode")
	}
}

func Init() {
	loadConfig()
	writeSyncer := getLogWriter()
	errWriteSyncer := getErrLogWriter()
	encoder := getEncoder()

	level := zap.AtomicLevel{}
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = zap.NewAtomicLevel()
	}
	var core zapcore.Core
	if appMode == appModeDebug {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, errWriteSyncer, zapcore.ErrorLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getErrLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   strings.TrimSuffix(filename, ".log") + ".err.log",
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
