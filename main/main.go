package main

import (
	"context"
	"fmt"
	pb "grpc-demo/proto"
	srv "grpc-demo/server"
	"log"
	"net"
)

const (
	grpcPort = ":6001"
	httpPort = ":8080"
)

type server struct{}

func (s *server) Hello(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Resp: "hello..."}, nil
}

func (s *server) Hello2(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Resp: "hello2..."}, nil
}
func main() {
	clientAddr := fmt.Sprintf("localhost%s", grpcPort)
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
			fmt.Printf("Failed to close %s %s: %v", "tcp", grpcPort, err)
		}
	}()

	go srv.StartHTTPServer(httpPort, clientAddr)
	log.Printf("GRPC Listening on %s\n", grpcPort)
	srv.StartGrpcServer(lis)
}
