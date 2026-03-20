package zap

import (
	"OperationAndMonitoring/config"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger

// InitLogger 初始化Logger
func InitLogger(cfg *config.Config) (err error) {
	atomicLevel, err := newAtomicLevel(cfg.Zap.Level, cfg.Zap.Debug)
	if err != nil {
		return err
	}

	writeSyncers := []zapcore.WriteSyncer{getLogWriter(cfg.Zap.FileName, cfg.Zap.MaxSize, cfg.Zap.MaxBackups, cfg.Zap.MaxAge)}
	if cfg.Zap.Debug {
		writeSyncers = append(writeSyncers, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(writeSyncers...),
		atomicLevel,
	)

	options := []zap.Option{zap.AddCaller()}
	if cfg.Zap.Debug {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	lg = zap.New(core, options...)
	zap.ReplaceGlobals(lg)
	return nil
}

func newAtomicLevel(level string, debugMode bool) (zap.AtomicLevel, error) {
	if debugMode && strings.TrimSpace(level) == "" {
		return zap.NewAtomicLevelAt(zap.DebugLevel), nil
	}
	if strings.TrimSpace(level) == "" {
		return zap.NewAtomicLevelAt(zap.InfoLevel), nil
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		return zap.AtomicLevel{}, fmt.Errorf("invalid zap.level %q: %w", level, err)
	}
	return zap.NewAtomicLevelAt(zapLevel), nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		CallerKey:      "caller",
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		message := c.Request.Method + " " + path
		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("cost", cost.String()),
			zap.String("ip", c.ClientIP()),
		}
		if query != "" {
			fields = append(fields, zap.String("query", query))
		}
		if privateErr := c.Errors.ByType(gin.ErrorTypePrivate).String(); privateErr != "" {
			fields = append(fields, zap.String("error", privateErr))
		}

		switch {
		case c.Writer.Status() >= http.StatusInternalServerError:
			lg.Error(message, fields...)
		case c.Writer.Status() >= http.StatusBadRequest:
			lg.Warn(message, fields...)
		case c.Writer.Status() != http.StatusNoContent:
			lg.Info(message, fields...)
		default:
			lg.Debug(message, fields...)
		}
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					lg.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					lg.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
