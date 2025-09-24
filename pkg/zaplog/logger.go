package zaplog

import (
	"os"
	"path"

	"github.com/mechiko/telebot_v4/internal/entity"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger
var EchoSugar *zap.Logger
var AuthSugar *zap.SugaredLogger

func init() {
	logpath := path.Join(entity.LogPath, "telebot.log")
	wLoggerRotate := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    5, // megabytes
		MaxBackups: 3,
		MaxAge:     1, // days
	})
	configLogger := zap.NewProductionEncoderConfig()
	configLogger.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// fileEncoder := zapcore.NewJSONEncoder(config)
	fileEncoder := zapcore.NewConsoleEncoder(configLogger)
	consoleEncoder := zapcore.NewConsoleEncoder(configLogger)
	// writer := zapcore.AddSync(logFile)
	writer := zapcore.AddSync(wLoggerRotate)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	// Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	Logger = zap.New(core, zap.AddCaller())

	configEcho := zap.NewProductionEncoderConfig()
	configEcho.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logpathecho := path.Join(entity.LogPath, "echo.log")
	echoLogFile, _ := os.OpenFile(logpathecho, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileEchoEncoder := zapcore.NewConsoleEncoder(configEcho)
	writerEcho := zapcore.AddSync(echoLogFile)
	core5 := zapcore.NewTee(
		zapcore.NewCore(fileEchoEncoder, writerEcho, defaultLogLevel),
	)
	EchoSugar = zap.New(core5, zap.AddCaller())

	configAuth := zap.NewProductionEncoderConfig()
	configAuth.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logpathAuth := path.Join(entity.LogPath, "auth.log")
	authLogFile, _ := os.OpenFile(logpathAuth, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileAuthEncoder := zapcore.NewConsoleEncoder(configEcho)
	writerAuth := zapcore.AddSync(authLogFile)
	coreAuth := zapcore.NewTee(
		zapcore.NewCore(fileAuthEncoder, writerAuth, defaultLogLevel),
	)
	AuthSugar = zap.New(coreAuth, zap.AddCaller()).Sugar()

}
