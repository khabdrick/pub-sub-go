package notifications

import (
	"context"
	"fmt"
	"time"

	"encore.dev/pubsub"
)

//encore:service
type Service struct{}

func initService() (*Service, error) {
	return &Service{}, nil
}

// UserSignupEvent represents a new user signup
type UserSignupEvent struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// OrderStatusEvent represents an order status change
type OrderStatusEvent struct {
	OrderID     string    `json:"order_id"`
	UserID      string    `json:"user_id"`
	OldStatus   string    `json:"old_status"`
	NewStatus   string    `json:"new_status"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ErrorEvent represents a system error
type ErrorEvent struct {
	ServiceName string    `json:"service_name"`
	ErrorCode   string    `json:"error_code"`
	Message     string    `json:"message"`
	Severity    string    `json:"severity"`
	Timestamp   time.Time `json:"timestamp"`
}

// Define topics
var (
	UserSignupTopic = pubsub.NewTopic[*UserSignupEvent]("user-signups", pubsub.TopicConfig{
		DeliveryGuarantee: pubsub.AtLeastOnce,
	})

	OrderStatusTopic = pubsub.NewTopic[*OrderStatusEvent]("order-status", pubsub.TopicConfig{
		DeliveryGuarantee: pubsub.AtLeastOnce,
	})

	ErrorTopic = pubsub.NewTopic[*ErrorEvent]("system-errors", pubsub.TopicConfig{
		DeliveryGuarantee: pubsub.AtLeastOnce,
	})
)

// Initialize subscribers
var _ = pubsub.NewSubscription(
	UserSignupTopic, "welcome-email",
	pubsub.SubscriptionConfig[*UserSignupEvent]{
		Handler: handleUserSignup,
	},
)

var _ = pubsub.NewSubscription(
	OrderStatusTopic, "order-status-notification",
	pubsub.SubscriptionConfig[*OrderStatusEvent]{
		Handler: handleOrderStatus,
	},
)

var _ = pubsub.NewSubscription(
	ErrorTopic, "error-alert",
	pubsub.SubscriptionConfig[*ErrorEvent]{
		Handler: handleError,
	},
)

// Subscriber handlers
func handleUserSignup(ctx context.Context, event *UserSignupEvent) error {
	// In a real implementation, you would integrate with an email service
	fmt.Printf("Sending welcome email to %s (UserID: %s)\n", event.Email, event.UserID)
	return nil
}

func handleOrderStatus(ctx context.Context, event *OrderStatusEvent) error {
	// In a real implementation, you would send push notifications or emails
	fmt.Printf("Order %s status changed from %s to %s for user %s\n",
		event.OrderID, event.OldStatus, event.NewStatus, event.UserID)
	return nil
}

func handleError(ctx context.Context, event *ErrorEvent) error {
	// In a real implementation, you would integrate with alert systems
	fmt.Printf("ALERT: %s error in %s: %s (Severity: %s)\n",
		event.ErrorCode, event.ServiceName, event.Message, event.Severity)
	return nil
}