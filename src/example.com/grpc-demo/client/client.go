package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "example.com/grpc-demo/pb/proto"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("Before call: %s, request: %+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("After call: %s, request: %+v", method, reply)
	return err
}

func main() {
	// WithInsecure without TLS connection
	// WithTransportCredentials with TLS connection
	// creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	// if err != nil {
	// 	log.Fatalf("Failed to certificate: %v", err)
	// }

	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptor))
	// conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithTransportCredentials(creds), grpc.WithUnaryInterceptor(unaryInterceptor))
	if err != nil {
		log.Fatalf("Client Connection Failed: %v", err)
	}
	// defer will be executed after main() is finished
	defer conn.Close()
	client := pb.NewSampleClient(conn)
	message := &pb.GetSampleMessage{Threshold: "default"}
	md := metadata.Pairs("timestamp", time.Now().Format(time.Stamp))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Once GetMyCat is executed, the thread will be stopped until the response will be returned.
	// Therefore, before executing GetMyCat, concurrently execute a function for canceling request with timeout.
	// This function is meant to cancel the request if the response will not be returned in a second.
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	res, err := client.GetSample(ctx, message, grpc.Trailer(&md))
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			log.Printf("gRPC Error (message: %s)", s.Message())
			for _, d := range s.Details() {
				switch info := d.(type) {
				case *errdetails.RetryInfo:
					log.Printf("RetryInfo: %v", info)
				}
			}
			os.Exit(1)
		} else {
			log.Fatalf("Could not get: %v", err)
		}
	}
	fmt.Printf("result:%#v \n", res)
}
