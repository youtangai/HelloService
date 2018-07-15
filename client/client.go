package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "github.com/youtangai/HelloStreaming/proto"
	"golang.org/x/net/context"
	"log"
	"time"
)

func runGreet(client pb.HelloServiceClient) {
	names := []string{"nagai", "nanaumi", "hiroto", "taichi", "miyoshi"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Greet(ctx)

	if err != nil {
		log.Fatalf("cannnot start greet: %v\n", err)
	}

	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Message:name}); err != nil {
			log.Fatalf("cannot send name: %v\n", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error occuerd when close: %v\n", err)
	}
	fmt.Println(reply)
}

func main() {
	fmt.Println("request to HelloServiceServer!!")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	runGreet(client)
}

