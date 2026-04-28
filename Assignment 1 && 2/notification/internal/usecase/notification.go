package usecase

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
)

type NotificationUsecase struct {
	Sender EmailPresenter
}

func (uc *NotificationUsecase) Send(ctx context.Context, notification *domain.Notification) error {
	return uc.Sender.Send(ctx, notification)
}
