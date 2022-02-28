package main

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	demo_middleware "github.com/shinshin8/golang-grpc-middleware/demo/server"
	pb "github.com/shinshin8/golang-grpc-protobuf/gen/go/protobuf"
	"github.com/shinshin8/golang-grpc-server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	srv := service.NewService()

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			demo_middleware.DemoServerInterceptor("one"),
			demo_middleware.DemoServerInterceptor("two"),
			demo_middleware.DemoServerInterceptor("three"),
			grpc_recovery.UnaryServerInterceptor()))

	pb.RegisterGrpcServiceServer(server, srv)

	lsn, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
