package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

const (
	appModeDebug  = "debug"
	logLevelInfo  = "info"
	logLevelError = "error"
)

var (
	appName    = "plugins"
	logLevel   = logLevelInfo
	logPath    = "/usr/local/var/log"
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
	level := zap.AtomicLevel{}
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = zap.NewAtomicLevel()
	}
	writeSyncer := logWriter(logLevelInfo)
	encoder := getEncoder()
	var core zapcore.Core
	if appMode == appModeDebug {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, logWriter(logLevelError), zapcore.ErrorLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
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

func logWriter(level string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPathWithAppNameAndLevel(level),
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func logPathWithAppNameAndLevel(level string) string {
	dir := path.Join(logPath, appName, level)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0666); err != nil {
			panic(err)
		}
	}
	return path.Join(dir, filename)
}
