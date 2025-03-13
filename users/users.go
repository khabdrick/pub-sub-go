package users

import (
	"context"
	"time"

	"encore.dev/beta/errs"
	"encore.app/notifications"
)

type SignupParams struct {
	Email    string `json:"email"`
	Password string `json:"password" encore:"sensitive"`
}

//encore:api public method=POST path=/users/signup
func Signup(ctx context.Context, params *SignupParams) error {
	// In a real implementation, you would store the user in a database
	userID := "user-123" // Generated user ID

	event := &notifications.UserSignupEvent{
		UserID:    userID,
		Email:     params.Email,
		CreatedAt: time.Now(),
	}

	if _, err := notifications.UserSignupTopic.Publish(ctx, event); err != nil {
		return errs.Wrap(err, "failed to publish signup event")
	}

	return nil
} 