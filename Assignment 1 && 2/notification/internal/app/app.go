package app

import (
	"github.com/fernoe1/AP2/assignment-1/notification/internal/adapter/cli"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/adapter/nats/notification/js"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/pkg/nats"
	"github.com/fernoe1/AP2/assignment-1/notification/internal/usecase"
)

func Start() {
	nc := nats.InitNATSConn()

	js.InitNotificationStream(nc)

	cliPresenter := cli.Presenter{}

	notificationUsecase := usecase.NotificationUsecase{
		Sender: cliPresenter,
	}

	consumer := js.NotificationConsumer{
		NotificationUsecase: notificationUsecase,
		Processed:           make(map[uint]bool),
	}

	consumer.ConsumeNotificationStream(nc)
}
