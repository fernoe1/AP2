package usecase

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
)

type EmailPresenter interface {
	Send(ctx context.Context, notification *domain.Notification) error
}
