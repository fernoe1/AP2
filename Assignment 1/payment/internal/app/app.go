package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	DB "github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/gorm"
	SERVER "github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/http/server"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/route"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/usecase"
	"github.com/fernoe1/AP2/assignment-1/payment/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	// database
	dsn := "host=localhost user=postgres password=130924 dbname=payment port=1987 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migration.Migrate(db)

	// repository
	paymentRepository := DB.PaymentRepository{Db: db}

	// usecase
	paymentUsecase := usecase.PaymentUsecase{PaymentRepository: &paymentRepository}

	// route
	r := route.InitRoute()
	route.RegisterPaymentRoute(r, &paymentUsecase)

	// server
	srv := SERVER.InitServer(":8082", r)

	start(&srv)
}

func start(srv *http.Server) {
	go func() {
		log.Println("payment starting at", srv.Addr)
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
