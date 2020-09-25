package main

import (
	"log"
	"net"

	pb "example.com/grpc-demo/pb/proto"
	"example.com/grpc-demo/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// cred, err := credentials.NewServerTLSFromFile("server.crt", "private.key")
	// if err != nil {
	// 	log.Fatalf("Failed to certificate: %v", err)
	// }

	server := grpc.NewServer()
	// server := grpc.NewServer(grpc.Creds(cred))

	service := &service.SampleService{}
	pb.RegisterSampleServer(server, service)
	reflection.Register(server)
	log.Printf("gRPC server listening on 19003")
	server.Serve(listenPort)
}
