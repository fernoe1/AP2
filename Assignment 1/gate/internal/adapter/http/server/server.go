package server

import (
	"net/http"
	"time"

	"github.com/fernoe1/AP2/assignment-1/gate/internal/route"
)

func InitServer(addr string) http.Server {
	return http.Server{
		Addr:              addr,
		Handler:           route.InitRoute(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
