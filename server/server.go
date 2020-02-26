package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-demo/proto"
	"log"
	"net"
	"net/http"
	"os"
)

type server struct{}

type errorBody struct {
	Err string `json:"error,omitempty"`
}

func (s *server) Hello(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Resp: "hello..."}, nil
}

func (s *server) Hello2(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Resp: "hello2..."}, nil
}

// start the http server
func StartHTTPServer(addr, clientAddr string) {
	log.Printf("Starting HTTP Server...")

	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()
	if err := pb.RegisterSayHelloHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
	log.Printf("HTTP Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// start grpc server
func StartGrpcServer(lis net.Listener) {
	log.Printf("Starting GRPC server...")
	s := grpc.NewServer()
	pb.RegisterSayHelloServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
		os.Exit(1)
	}
}
