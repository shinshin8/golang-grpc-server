package main

import (
	pb "github.com/shinshin8/golang-grpc-protobuf/gen/go/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	srv := NewService()

	server := grpc.NewServer(grpc.ChainUnaryInterceptor())

	pb.RegisterGrpcServiceServer(server, srv)

	lsn, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
