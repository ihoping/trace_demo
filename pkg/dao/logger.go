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
	*common.Logger
	LogLevel gormLogger.LogLevel
}

func NewLogger(ctx context.Context, logLevel gormLogger.LogLevel) *logger {
	traceData, ok := ctx.Value(common.TraceDataKey).(*common.TraceData)
	if !ok {
		traceData = &common.TraceData{}
	}

	commonLogger := common.NewLogger(common.LogLevel(logLevel), common.LogData{
		ProjectName: traceData.ProjectName,
		ClientID:    traceData.ClientID,
		TraceID:     traceData.TraceID,
		Action:      traceData.Action,
	})

	return &logger{
		Logger:   commonLogger,
		LogLevel: logLevel,
	}
}

// LogMode log mode
func (l *logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		l.Logger.Info("[GORM] " + msg + fmt.Sprintf("%v", data))
	}
}

// Warn print warn messages
func (l logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		l.Logger.Warn("[GORM] " + msg + fmt.Sprintf("%v", data))
	}
}

// Error print error messages
func (l logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		l.Logger.Error("[GORM] " + msg + fmt.Sprintf("%v", data))
	}
}

// Trace print sql message
func (l logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	l.Logger.Info("[GORM] " + elapsed.String() + " " + sql + " | " + strconv.Itoa(int(rows)) + " rows returned")
}
