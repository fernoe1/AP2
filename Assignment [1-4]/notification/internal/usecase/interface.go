package usecase

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
)

type EmailPresenter interface {
	Send(ctx context.Context, notification *domain.Notification) error
}

type NotificationStatusRepository interface {
	IsProcessed(ctx context.Context, paymentID string) (bool, error)
	MarkProcessed(ctx context.Context, paymentID string) error
	MarkFailed(ctx context.Context, paymentID string) error
}
