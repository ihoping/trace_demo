package common

import (
	"context"
	"errors"
)

const LoggerKey = "trace_logger"

func GetContext(logData LogData) context.Context {
	if logData.ProjectName == "" {
		logData.ProjectName = ProjectName
	}

	logger := NewLogger(LogInfo, logData)
	return context.WithValue(context.Background(), LoggerKey, logger)
}

func GetLogger(ctx context.Context) (*Logger, error) {
	logger, ok := ctx.Value(LoggerKey).(*Logger)
	if ok {
		return logger, nil
	}
	return nil, errors.New("logger not found")
}
