package main

import (
	"log"
	"net"

	pb "github.com/rkmangalp/go-grpc-basic/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloserver struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloserver{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
