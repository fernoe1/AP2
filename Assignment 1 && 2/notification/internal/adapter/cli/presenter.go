package cli

import (
	"context"
	"fmt"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
)

type Presenter struct {
}

func (c Presenter) Send(ctx context.Context, notification *domain.Notification) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	fmt.Printf("[Notification] Sent email to %s for Order #%d. Amount: $%d\n",
		notification.CustomerEmail, notification.ID, notification.Amount)

	return nil
}
