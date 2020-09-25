package service

import (
	"context"
	"log"
	"time"

	pb "example.com/grpc-demo/pb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SampleService struct{}

func (s *SampleService) GetSample(ctx context.Context, message *pb.GetSampleMessage) (*pb.ReturnResponse, error) {
	// grpc status codes
	// https://github.com/grpc/grpc/blob/master/doc/statuscodes.md

	log.Printf("Received: %v", message.Threshold)
	time.Sleep(3 * time.Second)

	switch message.Threshold {
	case "default":
		return &pb.ReturnResponse{
			Item: "default response",
		}, nil
	case "another":
		return &pb.ReturnResponse{
			Item: "another response",
		}, nil
	}
	// sample for not found error
	return nil, status.New(codes.NotFound, "Resource not found").Err()

	// sample for aborted error
	// st, _ := status.New(codes.Aborted, "aborted").WithDetails(&errdetails.RetryInfo{
	// 	RetryDelay: &duration.Duration{Seconds: 3, Nanos: 0},
	// })
	// return nil, st.Err()
}
