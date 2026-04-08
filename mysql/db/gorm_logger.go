package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
	logLevel logger.LogLevel
	enabled  bool
}

func NewGormLogger(level logger.LogLevel, enabled bool) logger.Interface {
	return &GormLogger{logLevel: level, enabled: enabled}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return &GormLogger{logLevel: level, enabled: l.enabled}
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel < logger.Info {
		return
	}
	if !l.enabled {
		return
	}
	log.Printf(msg, data...)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel < logger.Warn {
		return
	}
	if !l.enabled {
		return
	}
	log.Printf(msg, data...)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel < logger.Error {
		return
	}
	if !l.enabled {
		return
	}
	log.Printf(msg, data...)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if !l.enabled {
		return
	}
	if l.logLevel == logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()

	shouldLog := l.logLevel >= logger.Info || err != nil
	if !shouldLog {
		return
	}
	caller := "unknown caller"
	if v, ok := CallerFromContext(ctx); ok && v != "" {
		caller = v
	}

	log.Printf("\n%s %s\n[%0.3fms] [rows:%d] %s\n",
		time.Now().Format("2006/01/02 15:04:05"),
		caller,
		float64(elapsed.Microseconds())/1000.0,
		rows,
		sql,
	)
	if err != nil && err != os.ErrNotExist {
		log.Printf("gorm error: %s", fmt.Sprintf("%v", err))
	}
}
