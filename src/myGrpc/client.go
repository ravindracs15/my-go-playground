package main

import (
	"context"
	"log"
	"time"

	pb "my-go-playground/src/myGrpc/grpc_practice" // Replace with the actual path

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1000 * time.Second)
	defer cancel()

	for i := 0; i < 1000; i++ {
		response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World -- " + time.Now().String()})
		if err != nil {
			log.Fatalf("Could not greet: %v", err)
		}
		log.Printf("Greeting: %s", response.GetMessage())
		time.Sleep(time.Second)
	}
}
