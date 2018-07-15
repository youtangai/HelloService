package main

import (
	pb "github.com/youtangai/HelloStreaming/proto"
	"io"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
)

type helloService struct {}

func (hello *helloService) Greet(stream pb.HelloService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{
			    Message: "close connection",
			})
		}
		if err != nil {
			return err
		}
		fmt.Println("Hello " + req.Message + "!!")
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &helloService{})
	fmt.Println("start client-streaming type server!")
	grpcServer.Serve(lis)
}
