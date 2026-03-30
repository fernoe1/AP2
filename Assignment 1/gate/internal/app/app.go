package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernoe1/AP2/assignment-1/gate/internal/adapter/http/server"
)

type App struct {
	server http.Server
}

func InitApp(addr string) *App {
	return &App{server: server.InitServer(addr)}
}

func (a *App) Start() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	<-shutdownCh

	a.close()
}

func (a *App) close() {
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Println("server forced to shutdown:", err)
	}

	log.Println("server shutdown")
}
