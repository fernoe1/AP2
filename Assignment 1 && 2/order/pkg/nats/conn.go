package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func InitNATSConn() *nats.Conn {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal("failed to connect to NATS:", err)
	}

	return nc
}
