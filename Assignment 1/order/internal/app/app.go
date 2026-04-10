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

	DB "github.com/fernoe1/AP2/assignment-1/order/internal/adapter/gorm"
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/grpc"
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/grpc/service"
	HTTP "github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http"
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/route"
	"github.com/fernoe1/AP2/assignment-1/order/internal/usecase"
	"github.com/fernoe1/AP2/assignment-1/order/migration"
	GRPC "github.com/fernoe1/AP2/assignment-1/order/pkg/grpc"
	"github.com/fernoe1/AP2/assignment-1/order/pkg/nats"
	ordersvc "github.com/fernoe1/protogen/ap2-assign2/service/order"
	svc "github.com/fernoe1/protogen/ap2-assign2/service/payment"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	// database
	dsn := "host=order-db user=postgres password=130924 dbname=order port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migration.Migrate(db)

	// nats
	nc := nats.InitNATSConn()

	// repository
	orderRepository := DB.OrderRepository{Db: db, Nc: nc}

	// client
	// client := HTTP.InitClient()
	grpcConn, err := GRPC.InitGRPCConn(os.Getenv("GRPC_PAYMENT_SRV_ADDR"))
	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}

	grpcClient := grpc.InitClient(svc.NewPaymentServiceClient(grpcConn))

	// usecase
	orderUc := usecase.OrderUsecase{OrderRepository: &orderRepository, OrderClient: grpcClient}

	// route
	r := route.InitRoute()
	route.RegisterOrderRoute(r, &orderUc)

	// server
	srv := HTTP.InitServer(os.Getenv("HTTP_SRV_ADDR"), r)
	grpcSrv := grpc.InitServer()

	lis, err := net.Listen("tcp", os.Getenv("GRPC_SRV_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ordersvc.RegisterOrderServiceServer(grpcSrv, &service.OrderService{Nc: nc})

	reflection.Register(grpcSrv)

	go func() {
		log.Println("gRPC server listening on" + os.Getenv("GRPC_SRV_ADDR"))
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

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
