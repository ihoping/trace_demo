package dao

import (
	"context"
	"fmt"
	gormLogger "gorm.io/gorm/logger"
	"strconv"
	"time"
	"trace_demo/pkg/common"
)

type logger struct {
	LogLevel gormLogger.LogLevel
}

// LogMode log mode
func (l *logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l logger) Info(ctx context.Context, msg string, data ...interface{}) {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return
	}

	if l.LogLevel >= gormLogger.Info {
		commonLogger.Info("[GORM] " + msg + fmt.Sprintf("%v", data))
	}
}

// Warn print warn messages
func (l logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return
	}
	if l.LogLevel >= gormLogger.Warn {
		commonLogger.Warn("[GORM] " + msg + fmt.Sprintf("%v", data))
	}
}

// Error print error messages
func (l logger) Error(ctx context.Context, msg string, data ...interface{}) {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return
	}
	if l.LogLevel >= gormLogger.Error {
		commonLogger.Error("[GORM] " + msg + fmt.Sprintf("%v", data))
	}
}

// Trace print sql message
func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return
	}
	if l.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	commonLogger.Info("[GORM] " + elapsed.String() + " " + sql + " | " + strconv.Itoa(int(rows)) + " rows returned")
}
