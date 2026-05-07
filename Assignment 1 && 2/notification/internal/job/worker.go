package job

import (
	"context"
	"errors"
	"log"
	"math"
	"time"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/domain"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/usecase"
)

type Worker struct {
	notificationUsecase *usecase.NotificationUsecase
	maxRetries          int
	baseBackoff         time.Duration
	jobs                chan domain.Notification
}

func InitWorker(notificationUsecase *usecase.NotificationUsecase, maxRetries int, baseBackoff time.Duration, queueSize int) *Worker {
	return &Worker{
		notificationUsecase: notificationUsecase,
		maxRetries:          maxRetries,
		baseBackoff:         baseBackoff,
		jobs:                make(chan domain.Notification, queueSize),
	}
}

func (w *Worker) Enqueue(ctx context.Context, notification domain.Notification) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case w.jobs <- notification:
		return nil
	}
}

func (w *Worker) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case notification := <-w.jobs:
				if err := w.processWithRetry(ctx, notification); err != nil && !errors.Is(err, context.Canceled) {
					log.Printf("notification processing failed payment_id=%s: %v", notification.ID, err)
				}
			}
		}
	}()
}

func (w *Worker) processWithRetry(ctx context.Context, notification domain.Notification) error {
	var lastErr error

	for attempt := 0; attempt <= w.maxRetries; attempt++ {
		lastErr = w.notificationUsecase.Send(ctx, &notification)
		if lastErr == nil {
			return nil
		}

		if attempt == w.maxRetries {
			return lastErr
		}

		delay := w.baseBackoff * time.Duration(math.Pow(2, float64(attempt)))
		timer := time.NewTimer(delay)

		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}

	return lastErr
}
