package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/fernoe1/AP2/assignment-1/gate/internal/adapter/http"
	"github.com/fernoe1/AP2/assignment-1/gate/internal/route"
)

func Start() {
	// route
	r := route.InitRoute()

	// server
	srv := server.InitServer(":8080", r)

	start(&srv)
}

func start(srv *http.Server) {
	go func() {
		log.Println("gate starting at", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	<-shutdownCh

	shutdown(srv)
}

func shutdown(srv *http.Server) {
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("server forced to shutdown:", err)
	}

	log.Println("server shutdown")
}
