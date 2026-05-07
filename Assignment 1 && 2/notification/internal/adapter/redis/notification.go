package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	statusProcessed = "processed"
	statusFailed    = "failed"
)

type NotificationRepository struct {
	client *redis.Client
	ttl    time.Duration
}

func NewIdempotencyStore(client *redis.Client, ttl time.Duration) *NotificationRepository {
	return &NotificationRepository{client: client, ttl: ttl}
}

func (s *NotificationRepository) IsProcessed(ctx context.Context, paymentID string) (bool, error) {
	value, err := s.client.Get(ctx, s.key(paymentID)).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return value == statusProcessed, nil
}

func (s *NotificationRepository) MarkProcessed(ctx context.Context, paymentID string) error {
	return s.client.Set(ctx, s.key(paymentID), statusProcessed, s.ttl).Err()
}

func (s *NotificationRepository) MarkFailed(ctx context.Context, paymentID string) error {
	return s.client.Set(ctx, s.key(paymentID), statusFailed, s.ttl).Err()
}

func (s *NotificationRepository) key(paymentID string) string {
	return "notification:payment:" + paymentID
}
