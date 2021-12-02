package common

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

var globalWriter io.Writer = os.Stdout //默认标准输出

func SetGlobalWriter(writer io.Writer) {
	globalWriter = writer
}

type Logger struct {
	io.Writer
	Context     context.Context
	LogLevel    LogLevel
	projectName string
	clientID    string
	traceID     string
	action      string
}

type LogLevel uint8

const (
	LogSilent LogLevel = iota + 1
	LogError
	LogWarn
	LogInfo
)

type LogData struct {
	ProjectName string
	ClientID    string
	TraceID     string
	Action      string
}

func NewLogger(logLevel LogLevel, data LogData) *Logger {
	return &Logger{
		Writer:      globalWriter,
		LogLevel:    logLevel,
		projectName: data.ProjectName,
		clientID:    data.ClientID,
		traceID:     data.TraceID,
		action:      data.Action,
	}
}

// Info print info
func (l Logger) Info(msg string) {
	if l.LogLevel >= LogInfo {
		l.Printf(LogInfo, msg)
	}
}

// Warn print error messages
func (l Logger) Warn(msg string) {
	if l.LogLevel >= LogWarn {
		l.Printf(LogWarn, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.LogLevel >= LogError {
		l.Printf(LogError, msg)
	}
}

func (l Logger) Printf(level LogLevel, content string) {
	//time project_name level action content @trace_id
	//2021-09-10 12:30:21 trace_demo 1 blog/get-detail [GORM]select * from t_article limit 1  @AX1vYm-DSQmUAAnr
	msg := fmt.Sprintf("%s %s %d %s %s @%s-%s\n", time.Now().UTC().Format(time.RFC3339), l.projectName, level, l.action, content, l.clientID, l.traceID)

	_, _ = l.Writer.Write([]byte(msg))
}
