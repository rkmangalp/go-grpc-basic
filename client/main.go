package main

import (
	"log"

	pb "github.com/rkmangalp/go-grpc-basic/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Veena", "Ravi", "kiran"},
	}

	// callSayHello(client)
	// callSayHello(client)

	//call server streamning
	// callSayHelloServerStream(client, names)

	// Call client streaming
	// callSayHelloClientStream(client, names)

	//call bidirectional stream
	callHellobidirectionalstream(client, names)
}
