package orders

import (
	"context"
	"time"

	"encore.dev/beta/errs"
	"encore.app/notifications"
)

type UpdateOrderParams struct {
	OrderID   string `json:"order_id"`
	UserID    string `json:"user_id"`
	NewStatus string `json:"new_status"`
}

//encore:api public method=POST path=/orders/update-status
func UpdateOrderStatus(ctx context.Context, params *UpdateOrderParams) error {
	// In a real implementation, you would update the order in a database
	oldStatus := "Pending" // This would come from the database

	event := &notifications.OrderStatusEvent{
		OrderID:   params.OrderID,
		UserID:    params.UserID,
		OldStatus: oldStatus,
		NewStatus: params.NewStatus,
		UpdatedAt: time.Now(),
	}

	if _, err := notifications.OrderStatusTopic.Publish(ctx, event); err != nil {
		return errs.Wrap(err, "failed to publish order status event")
	}

	return nil
} 