package app

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/fernoe1/AP2/assignment-1/notification/internal/adapter/mailjet"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/adapter/nats/notification/js"
	redisadapter "github.com/fernoe1/AP2/assignment-1/notification/internal/adapter/redis"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/job"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/pkg/nats"
	redispkg "github.com/fernoe1/AP2/assignment-1/notification/internal/pkg/redis"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/usecase"
)

func Start() {
	ctx := context.Background()

	nc := nats.InitNATSConn()
	redisClient := redispkg.InitRDBConn()

	js.InitNotificationStream(nc)

	emailProvider := mailjet.InitClient()
	idempotencyStore := redisadapter.NewIdempotencyStore(redisClient, 5*time.Minute)

	notificationUsecase := usecase.NotificationUsecase{
		Sender: emailProvider,
		Repo:   idempotencyStore,
	}

	maxRetries, _ := strconv.Atoi(os.Getenv("MAX_RETRIES"))
	ttl, _ := strconv.Atoi(os.Getenv("TTL"))

	worker := job.InitWorker(&notificationUsecase, maxRetries, time.Second*time.Duration(ttl), 10)
	worker.Start(ctx)

	consumer := js.NotificationConsumer{
		Worker: worker,
	}

	consumer.ConsumeNotificationStream(nc)
}
