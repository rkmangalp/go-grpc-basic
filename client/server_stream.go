package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rkmangalp/go-grpc-basic/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	// Log the start of the streaming process
	log.Println("Streaming started")

	// Initiate a server-streaming RPC call to the SayHelloServerStreaming method with the provided context and names
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Failed to send names: %v", err)
	}

	// Continuously receive messages from the stream until an error or EOF (end of file) is encountered
	for {
		// Receive a message from the stream
		message, err := stream.Recv()
		if err == io.EOF {
			// Exit the loop when the stream ends (EOF)
			break
		}
		if err != nil {
			// Log and exit if any other error occurs during streaming
			log.Fatalf("Error while streaming: %v", err)
		}
		// Log the received message
		log.Println("Received message:", message)
	}

	// Log the end of the streaming process
	log.Println("Streaming finished")
}
