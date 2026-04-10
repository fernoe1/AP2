package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	DB "github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/gorm"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/grpc"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/grpc/service"
	SERVER "github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/http"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/http/route"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/usecase"
	"github.com/fernoe1/AP2/assignment-1/payment/migration"
	paymentsvc "github.com/fernoe1/protogen/ap2-assign2/service/payment"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	// database
	dsn := "host=payment-db user=postgres password=130924 dbname=payment port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
	route.RegisterPaymentRoute(r, paymentUsecase)

	// server
	srv := SERVER.InitServer(":8082", r)
	grpcSrv := grpc.InitServer()

	lis, err := net.Listen("tcp", os.Getenv("GRPC_SRV_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	paymentsvc.RegisterPaymentServiceServer(grpcSrv, &service.PaymentService{UC: &paymentUsecase})

	reflection.Register(grpcSrv)

	log.Println("gRPC server listening on" + os.Getenv("GRPC_SRV_ADDR"))
	go func() {
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

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
