package usecase

import (
	"context"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
)

type NotificationUsecase struct {
	Sender EmailPresenter
	Repo   NotificationStatusRepository
}

func (uc *NotificationUsecase) Send(ctx context.Context, notification *domain.Notification) error {
	processed, err := uc.Repo.IsProcessed(ctx, notification.ID)
	if err != nil {
		return err
	}
	if processed {
		return nil
	}

	if err := uc.Sender.Send(ctx, notification); err != nil {
		_ = uc.Repo.MarkFailed(ctx, notification.ID)
		return err
	}

	return uc.Repo.MarkProcessed(ctx, notification.ID)
}
