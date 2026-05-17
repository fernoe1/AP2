package grpc

import "google.golang.org/grpc"

func InitServer() *grpc.Server {
	return grpc.NewServer()
}
