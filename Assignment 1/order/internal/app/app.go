package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gormDB "github.com/fernoe1/AP2/assignment-1/order/internal/adapter/gorm"
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/server"
	"github.com/fernoe1/AP2/assignment-1/order/internal/route"
	"github.com/fernoe1/AP2/assignment-1/order/internal/usecase"
	"github.com/fernoe1/AP2/assignment-1/order/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	// database
	dsn := "host=localhost user=postgres password=130924 dbname=gorm port=1987 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migration.Migrate(db)

	// repository
	orderRepository := gormDB.OrderRepository{Db: db}

	// usecase
	orderUc := usecase.OrderUsecase{OrderRepository: &orderRepository}

	// route
	r := route.InitRoute()
	route.RegisterOrderRoute(r, orderUc)

	// server
	srv := server.InitServer(":8081", r)

	start(&srv)
}

func start(srv *http.Server) {
	go func() {
		log.Println("order starting at", srv.Addr)
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
