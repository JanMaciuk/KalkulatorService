package main

import (
	"fmt"
	"github/JahoPL/CalculatorService/files"
	"github/JahoPL/CalculatorService/server/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	port := "50051"
	lis, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	greetService := server.NewCaluclatorService()
	//pb.RegisterRouteGuideServer(grpcServer, newServer())
	files.RegisterGreetServiceServer(grpcServer, greetService)
	fmt.Println("Server running ....")
	grpcServer.Serve(lis)
}
