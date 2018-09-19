package main

import (
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"go-grpc-app/service"
	"go-grpc-app/pb"
	"google.golang.org/grpc"
)

const PORT string = ":5000"

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &service.UserService{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
