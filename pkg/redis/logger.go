package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"trace_demo/pkg/common"
)

type hook struct {
}

func (hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return ctx, nil
	}
	commonLogger.Info(fmt.Sprintf("[redis] %s", cmd))
	return ctx, nil
}

func (hook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return nil
	}
	commonLogger.Info(fmt.Sprintf("[redis] %s", cmd))
	return nil
}

func (hook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return ctx, nil
	}
	commonLogger.Info(fmt.Sprintf("[redis] %v", cmds))
	return ctx, nil
}

func (hook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	commonLogger, err := common.GetLogger(ctx)
	if err != nil {
		return nil
	}
	commonLogger.Info(fmt.Sprintf("[redis] %v", cmds))
	return nil
}
