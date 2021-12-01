package common

import (
	"context"
)

type TraceData struct {
	ProjectName string
	ClientID    string
	TraceID     string
	Action      string
}

const TraceDataKey = "trace_data"

func GetContext(traceData *TraceData) context.Context {
	if traceData.ProjectName == "" {
		traceData.ProjectName = ProjectName
	}

	return context.WithValue(context.Background(), TraceDataKey, traceData)
}
