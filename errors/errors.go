package errors

import (
	"context"
	"time"

	"encore.dev/beta/errs"
	"encore.app/notifications"
)

type LogErrorParams struct {
	ServiceName string `json:"service_name"`
	ErrorCode   string `json:"error_code"`
	Message     string `json:"message"`
	Severity    string `json:"severity"`
}

//encore:api public method=POST path=/errors/log
func LogError(ctx context.Context, params *LogErrorParams) error {
	event := &notifications.ErrorEvent{
		ServiceName: params.ServiceName,
		ErrorCode:   params.ErrorCode,
		Message:     params.Message,
		Severity:    params.Severity,
		Timestamp:   time.Now(),
	}

	if _, err := notifications.ErrorTopic.Publish(ctx, event); err != nil {
		return errs.Wrap(err, "failed to publish error event")
	}

	return nil
} 